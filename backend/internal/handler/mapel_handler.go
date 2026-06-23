package handler

import (
	"net/http"
	"strconv"

	"edu-admin/internal/database"
	"edu-admin/internal/dto"
	"edu-admin/internal/model"

	"github.com/gin-gonic/gin"
)

func GetMapel(c *gin.Context) {
	var list []model.MataPelajaran
	database.DB.Order("created_at desc").Find(&list)
	c.JSON(http.StatusOK, dto.OK(list))
}

func CreateMapel(c *gin.Context) {
	var req dto.CreateMapelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error("Input tidak valid: "+err.Error()))
		return
	}

	m := model.MataPelajaran{Kode: req.Kode, Nama: req.Nama, KKM: req.KKM}
	if err := database.DB.Create(&m).Error; err != nil {
		c.JSON(http.StatusBadRequest, dto.Error("Gagal membuat mapel (kode mungkin sudah ada)"))
		return
	}
	c.JSON(http.StatusCreated, dto.Created(m))
}

func UpdateMapel(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var m model.MataPelajaran
	if err := database.DB.First(&m, uint(id)).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.Error("Mapel tidak ditemukan"))
		return
	}

	var req dto.UpdateMapelRequest
	c.ShouldBindJSON(&req)

	updates := map[string]interface{}{}
	if req.Kode != "" {
		updates["kode"] = req.Kode
	}
	if req.Nama != "" {
		updates["nama"] = req.Nama
	}
	if req.KKM != 0 {
		updates["kkm"] = req.KKM
	}
	database.DB.Model(&m).Updates(updates)
	database.DB.First(&m, uint(id))
	c.JSON(http.StatusOK, dto.OK(m))
}

func DeleteMapel(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	database.DB.Delete(&model.MataPelajaran{}, uint(id))
	c.JSON(http.StatusOK, dto.OK(nil))
}
