package handler

import (
	"net/http"
	"strconv"

	"edu-admin/internal/database"
	"edu-admin/internal/dto"
	"edu-admin/internal/model"

	"github.com/gin-gonic/gin"
)

func GetKelas(c *gin.Context) {
	query := database.DB.Preload("TahunAjaran").Order("created_at desc")

	if t := c.Query("tingkat"); t != "" {
		query = query.Where("tingkat = ?", t)
	}
	if ta := c.Query("tahun_ajaran_id"); ta != "" {
		query = query.Where("tahun_ajaran_id = ?", ta)
	}

	var list []model.Kelas
	query.Find(&list)
	c.JSON(http.StatusOK, dto.OK(list))
}

func CreateKelas(c *gin.Context) {
	var req dto.CreateKelasRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error("Input tidak valid: "+err.Error()))
		return
	}

	k := model.Kelas{
		Nama: req.Nama, Tingkat: req.Tingkat,
		TahunAjaranID: req.TahunAjaranID, WaliKelas: req.WaliKelas,
	}
	if err := database.DB.Create(&k).Error; err != nil {
		c.JSON(http.StatusBadRequest, dto.Error("Gagal membuat kelas"))
		return
	}
	database.DB.Preload("TahunAjaran").First(&k, k.ID)
	c.JSON(http.StatusCreated, dto.Created(k))
}

func UpdateKelas(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var k model.Kelas
	if err := database.DB.First(&k, uint(id)).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.Error("Kelas tidak ditemukan"))
		return
	}

	var req dto.UpdateKelasRequest
	c.ShouldBindJSON(&req)

	updates := map[string]interface{}{}
	if req.Nama != "" {
		updates["nama"] = req.Nama
	}
	if req.Tingkat != "" {
		updates["tingkat"] = req.Tingkat
	}
	if req.TahunAjaranID != 0 {
		updates["tahun_ajaran_id"] = req.TahunAjaranID
	}
	if req.WaliKelas != "" {
		updates["wali_kelas"] = req.WaliKelas
	}
	database.DB.Model(&k).Updates(updates)
	database.DB.Preload("TahunAjaran").First(&k, uint(id))
	c.JSON(http.StatusOK, dto.OK(k))
}

func DeleteKelas(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	database.DB.Delete(&model.Kelas{}, uint(id))
	c.JSON(http.StatusOK, dto.OK(nil))
}
