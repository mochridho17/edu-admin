package handler

import (
	"net/http"
	"strconv"

	"edu-admin/internal/database"
	"edu-admin/internal/dto"
	"edu-admin/internal/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *gin.Context) {
	var users []model.User
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error("Gagal mengambil data user"))
		return
	}

	var result []dto.UserResponse
	for _, u := range users {
		result = append(result, dto.UserResponse{
			ID: u.ID, Username: u.Username, Email: u.Email, Role: u.Role,
		})
	}
	c.JSON(http.StatusOK, dto.OK(result))
}

func CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error("Input tidak valid: "+err.Error()))
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error("Gagal hash password"))
		return
	}

	user := model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hash),
		Role:    req.Role,
	}
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, dto.Error("Gagal membuat user: mungkin username/email sudah ada"))
		return
	}

	c.JSON(http.StatusCreated, dto.Created(dto.UserResponse{
		ID: user.ID, Username: user.Username, Email: user.Email, Role: user.Role,
	}))
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	var user model.User
	if err := database.DB.First(&user, uint(id)).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.Error("User tidak ditemukan"))
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error("Input tidak valid: "+err.Error()))
		return
	}

	updates := map[string]interface{}{}
	if req.Username != "" {
		updates["username"] = req.Username
	}
	if req.Email != "" {
		updates["email"] = req.Email
	}
	if req.Role != "" {
		updates["role"] = req.Role
	}
	if req.Password != "" {
		hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		updates["password"] = string(hash)
	}

	if err := database.DB.Model(&user).Updates(updates).Error; err != nil {
		c.JSON(http.StatusBadRequest, dto.Error("Gagal update user"))
		return
	}

	database.DB.First(&user, uint(id))
	c.JSON(http.StatusOK, dto.OK(dto.UserResponse{
		ID: user.ID, Username: user.Username, Email: user.Email, Role: user.Role,
	}))
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := database.DB.Delete(&model.User{}, uint(id)).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error("Gagal hapus user"))
		return
	}
	c.JSON(http.StatusOK, dto.OK(nil))
}
