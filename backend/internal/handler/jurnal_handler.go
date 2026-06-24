package handler

import (
	"net/http"
	"strconv"

	"edu-admin/internal/database"
	"edu-admin/internal/dto"
	"edu-admin/internal/model"

	"github.com/gin-gonic/gin"
)

func GetJurnalGuru(c *gin.Context) {
	query := database.DB.Preload("Mapel").Preload("Kelas").Preload("Kelas.TahunAjaran").Preload("TahunAjaran").Order("tanggal desc")

	if nama := c.Query("nama_guru"); nama != "" {
		query = query.Where("nama_guru ILIKE ?", "%"+nama+"%")
	}
	if mapelID := c.Query("mapel_id"); mapelID != "" {
		query = query.Where("mapel_id = ?", mapelID)
	}
	if kelasID := c.Query("kelas_id"); kelasID != "" {
		query = query.Where("kelas_id = ?", kelasID)
	}
	if semester := c.Query("semester"); semester != "" {
		query = query.Where("semester = ?", semester)
	}
	if taID := c.Query("tahun_ajaran_id"); taID != "" {
		query = query.Where("tahun_ajaran_id = ?", taID)
	}

	var list []model.JurnalGuru
	query.Find(&list)
	c.JSON(http.StatusOK, dto.OK(list))
}

func CreateJurnalGuru(c *gin.Context) {
	var req dto.CreateJurnalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error("Input tidak valid: "+err.Error()))
		return
	}

	j := model.JurnalGuru{
		Tanggal:       req.Tanggal,
		NamaGuru:      req.NamaGuru,
		MapelID:       req.MapelID,
		KelasID:       req.KelasID,
		Kegiatan:      req.Kegiatan,
		Semester:      req.Semester,
		TahunAjaranID: req.TahunAjaranID,
	}
	if err := database.DB.Create(&j).Error; err != nil {
		c.JSON(http.StatusBadRequest, dto.Error("Gagal membuat jurnal"))
		return
	}
	database.DB.Preload("Mapel").Preload("Kelas").Preload("TahunAjaran").First(&j, j.ID)
	c.JSON(http.StatusCreated, dto.Created(j))
}

func UpdateJurnalGuru(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var j model.JurnalGuru
	if err := database.DB.First(&j, uint(id)).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.Error("Jurnal tidak ditemukan"))
		return
	}

	var req dto.UpdateJurnalRequest
	c.ShouldBindJSON(&req)

	updates := map[string]interface{}{}
	if req.Tanggal != "" {
		updates["tanggal"] = req.Tanggal
	}
	if req.NamaGuru != "" {
		updates["nama_guru"] = req.NamaGuru
	}
	if req.MapelID != 0 {
		updates["mapel_id"] = req.MapelID
	}
	if req.KelasID != 0 {
		updates["kelas_id"] = req.KelasID
	}
	if req.Kegiatan != "" {
		updates["kegiatan"] = req.Kegiatan
	}
	if req.Semester != "" {
		updates["semester"] = req.Semester
	}
	if req.TahunAjaranID != 0 {
		updates["tahun_ajaran_id"] = req.TahunAjaranID
	}
	database.DB.Model(&j).Updates(updates)
	database.DB.Preload("Mapel").Preload("Kelas").Preload("TahunAjaran").First(&j, uint(id))
	c.JSON(http.StatusOK, dto.OK(j))
}

func DeleteJurnalGuru(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	database.DB.Delete(&model.JurnalGuru{}, uint(id))
	c.JSON(http.StatusOK, dto.OK(nil))
}
