package handler

import (
	"math"
	"net/http"
	"strconv"

	"edu-admin/internal/database"
	"edu-admin/internal/dto"
	"edu-admin/internal/model"

	"github.com/gin-gonic/gin"
)

func GetNilai(c *gin.Context) {
	kelasID := c.Query("kelas_id")
	mapelID := c.Query("mapel_id")
	semester := c.Query("semester")
	tahunAjaranID := c.Query("tahun_ajaran_id")

	query := database.DB.
		Preload("Siswa").
		Preload("Mapel").
		Preload("TahunAjaran")

	if kelasID != "" {
		query = query.Joins("JOIN siswa ON siswa.id = nilai.siswa_id").
			Where("siswa.kelas_id = ?", kelasID)
	}
	if mapelID != "" {
		query = query.Where("nilai.mapel_id = ?", mapelID)
	}
	if semester != "" {
		query = query.Where("nilai.semester = ?", semester)
	}
	if tahunAjaranID != "" {
		query = query.Where("nilai.tahun_ajaran_id = ?", tahunAjaranID)
	}

	var list []model.Nilai
	query.Order("siswa.nama asc").Find(&list)
	c.JSON(http.StatusOK, dto.OK(list))
}

func CreateNilaiBatch(c *gin.Context) {
	var req dto.CreateNilaiBatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error("Input tidak valid: "+err.Error()))
		return
	}

	for _, item := range req.Data {
		// Compute akhir
		tugas := 0.0
		uts := 0.0
		uas := 0.0
		if item.Tugas != nil {
			tugas = *item.Tugas
		}
		if item.UTS != nil {
			uts = *item.UTS
		}
		if item.UAS != nil {
			uas = *item.UAS
		}

		akhir := math.Round(tugas*0.2+uts*0.3+uas*0.5) / 100 * 100

		// Upsert
		var existing model.Nilai
		result := database.DB.Where("siswa_id = ? AND mapel_id = ? AND semester = ? AND tahun_ajaran_id = ?",
			item.SiswaID, req.MapelID, req.Semester, req.TahunAjaranID).
			First(&existing)

		if result.Error != nil {
			// Create
			n := model.Nilai{
				SiswaID:       item.SiswaID,
				MapelID:       req.MapelID,
				Semester:      req.Semester,
				TahunAjaranID: req.TahunAjaranID,
				Tugas:         item.Tugas,
				UTS:           item.UTS,
				UAS:           item.UAS,
				Akhir:         &akhir,
			}
			database.DB.Create(&n)
		} else {
			// Update
			updates := map[string]interface{}{
				"tugas": item.Tugas, "uts": item.UTS, "uas": item.UAS, "akhir": akhir,
			}
			database.DB.Model(&existing).Updates(updates)
		}
	}

	// Return updated data
	GetNilai(c)
}

func RekapNilaiSiswa(c *gin.Context) {
	siswaID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	var nilai []model.Nilai
	database.DB.Preload("Mapel").
		Where("siswa_id = ?", uint(siswaID)).
		Find(&nilai)

	var result []dto.NilaiRekapSiswa
	for _, n := range nilai {
		if n.Mapel == nil {
			continue
		}
		tuntas := false
		if n.Akhir != nil {
			tuntas = *n.Akhir >= float64(n.Mapel.KKM)
		}
		result = append(result, dto.NilaiRekapSiswa{
			MapelNama: n.Mapel.Nama,
			MapelKode: n.Mapel.Kode,
			KKM:       n.Mapel.KKM,
			Tugas:     n.Tugas,
			UTS:       n.UTS,
			UAS:       n.UAS,
			Akhir:     n.Akhir,
			Tuntas:    tuntas,
		})
	}

	c.JSON(http.StatusOK, dto.OK(result))
}

func RankingKelas(c *gin.Context) {
	kelasID := c.Query("kelas_id")
	semester := c.Query("semester")
	tahunAjaranID := c.Query("tahun_ajaran_id")

	type RataSiswa struct {
		SiswaID  uint
		Nama     string
		NIS      string
		RataRata float64
	}

	var results []RataSiswa
	database.DB.Table("nilai").
		Select("siswa.id as siswa_id, siswa.nama, siswa.nis, AVG(nilai.akhir) as rata_rata").
		Joins("JOIN siswa ON siswa.id = nilai.siswa_id").
		Where("siswa.kelas_id = ?", kelasID).
		Where("nilai.semester = ?", semester).
		Where("nilai.tahun_ajaran_id = ?", tahunAjaranID).
		Group("siswa.id, siswa.nama, siswa.nis").
		Order("rata_rata desc").
		Scan(&results)

	var ranking []dto.NilaiRankingKelas
	for i, r := range results {
		ranking = append(ranking, dto.NilaiRankingKelas{
			SiswaID:   r.SiswaID,
			NamaSiswa: r.Nama,
			NIS:       r.NIS,
			RataRata:  math.Round(r.RataRata*10) / 10,
			Peringkat: i + 1,
		})
	}

	c.JSON(http.StatusOK, dto.OK(ranking))
}
