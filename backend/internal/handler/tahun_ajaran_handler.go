package handler

import (
	"net/http"
	"strconv"

	"edu-admin/internal/database"
	"edu-admin/internal/dto"
	"edu-admin/internal/model"

	"github.com/gin-gonic/gin"
)

func GetTahunAjaran(c *gin.Context) {
	var list []model.TahunAjaran
	database.DB.Order("created_at desc").Find(&list)
	c.JSON(http.StatusOK, dto.OK(list))
}

func CreateTahunAjaran(c *gin.Context) {
	var req dto.CreateTahunAjaranRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error("Input tidak valid: "+err.Error()))
		return
	}

	ta := model.TahunAjaran{Nama: req.Nama, Semester: req.Semester}
	if err := database.DB.Create(&ta).Error; err != nil {
		c.JSON(http.StatusBadRequest, dto.Error("Gagal membuat tahun ajaran"))
		return
	}
	c.JSON(http.StatusCreated, dto.Created(ta))
}

func UpdateTahunAjaran(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var ta model.TahunAjaran
	if err := database.DB.First(&ta, uint(id)).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.Error("Tahun ajaran tidak ditemukan"))
		return
	}

	var req dto.UpdateTahunAjaranRequest
	c.ShouldBindJSON(&req)

	updates := map[string]interface{}{}
	if req.Nama != "" {
		updates["nama"] = req.Nama
	}
	if req.Semester != "" {
		updates["semester"] = req.Semester
	}
	database.DB.Model(&ta).Updates(updates)
	database.DB.First(&ta, uint(id))
	c.JSON(http.StatusOK, dto.OK(ta))
}

func DeleteTahunAjaran(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	database.DB.Delete(&model.TahunAjaran{}, uint(id))
	c.JSON(http.StatusOK, dto.OK(nil))
}

func ActivateTahunAjaran(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	// Deactivate all
	database.DB.Model(&model.TahunAjaran{}).Update("is_active", false)

	// Activate selected
	database.DB.Model(&model.TahunAjaran{}).Where("id = ?", uint(id)).Update("is_active", true)

	var ta model.TahunAjaran
	database.DB.First(&ta, uint(id))
	c.JSON(http.StatusOK, dto.OK(ta))
}
