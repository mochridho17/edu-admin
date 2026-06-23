package handler

import (
	"fmt"
	"net/http"
	"time"

	"edu-admin/internal/database"
	"edu-admin/internal/dto"
	"edu-admin/internal/model"

	"github.com/gin-gonic/gin"
)

func GetDashboardStats(c *gin.Context) {
	var totalSiswa, totalKelas, totalMapel int64
	database.DB.Model(&model.Siswa{}).Count(&totalSiswa)
	database.DB.Model(&model.Kelas{}).Count(&totalKelas)
	database.DB.Model(&model.MataPelajaran{}).Count(&totalMapel)

	type KehadiranCount struct {
		Status string
		Total  int64
	}
	var kehadiran []KehadiranCount
	tanggalHariIni := time.Now().Format("2006-01-02")

	database.DB.Model(&model.Absensi{}).
		Select("status, COUNT(*) as total").
		Where("tanggal = ?", tanggalHariIni).
		Group("status").
		Find(&kehadiran)

	totalKehadiran := int64(0)
	totalHadir := int64(0)
	absensiHariIni := map[string]int64{"hadir": 0, "sakit": 0, "izin": 0, "alfa": 0}

	for _, k := range kehadiran {
		absensiHariIni[k.Status] = k.Total
		totalKehadiran += k.Total
		if k.Status == "hadir" {
			totalHadir = k.Total
		}
	}

	rataKehadiran := "0%"
	if totalKehadiran > 0 {
		persen := float64(totalHadir) / float64(totalKehadiran) * 100
		rataKehadiran = fmt.Sprintf("%.1f%%", persen)
	}

	c.JSON(http.StatusOK, dto.OK(dto.DashboardStats{
		TotalSiswa:        totalSiswa,
		TotalKelas:        totalKelas,
		TotalMapel:        totalMapel,
		RataRataKehadiran: rataKehadiran,
		AbsensiHariIni:    absensiHariIni,
	}))
}
