package handler

import (
	"net/http"
	"strconv"

	"edu-admin/internal/database"
	"edu-admin/internal/dto"
	"edu-admin/internal/model"

	"github.com/gin-gonic/gin"
)

func GetSiswa(c *gin.Context) {
	query := database.DB.Preload("Kelas").Order("created_at desc")

	if k := c.Query("kelas_id"); k != "" {
		query = query.Where("kelas_id = ?", k)
	}
	if s := c.Query("search"); s != "" {
		query = query.Where("nis ILIKE ? OR nama ILIKE ?", "%"+s+"%", "%"+s+"%")
	}

	var list []model.Siswa
	query.Find(&list)
	c.JSON(http.StatusOK, dto.OK(list))
}

func CreateSiswa(c *gin.Context) {
	var req dto.CreateSiswaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error("Input tidak valid: "+err.Error()))
		return
	}

	s := model.Siswa{
		NIS: req.NIS, Nama: req.Nama,
		JenisKelamin: req.JenisKelamin, KelasID: req.KelasID,
	}
	if err := database.DB.Create(&s).Error; err != nil {
		c.JSON(http.StatusBadRequest, dto.Error("Gagal membuat siswa (NIS mungkin sudah ada)"))
		return
	}
	database.DB.Preload("Kelas").First(&s, s.ID)
	c.JSON(http.StatusCreated, dto.Created(s))
}

func UpdateSiswa(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var s model.Siswa
	if err := database.DB.First(&s, uint(id)).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.Error("Siswa tidak ditemukan"))
		return
	}

	var req dto.UpdateSiswaRequest
	c.ShouldBindJSON(&req)

	updates := map[string]interface{}{}
	if req.NIS != "" {
		updates["nis"] = req.NIS
	}
	if req.Nama != "" {
		updates["nama"] = req.Nama
	}
	if req.JenisKelamin != "" {
		updates["jenis_kelamin"] = req.JenisKelamin
	}
	if req.KelasID != 0 {
		updates["kelas_id"] = req.KelasID
	}
	database.DB.Model(&s).Updates(updates)
	database.DB.Preload("Kelas").First(&s, uint(id))
	c.JSON(http.StatusOK, dto.OK(s))
}

func DeleteSiswa(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	database.DB.Delete(&model.Siswa{}, uint(id))
	c.JSON(http.StatusOK, dto.OK(nil))
}
