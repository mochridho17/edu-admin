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

func GetAbsensi(c *gin.Context) {
	kelasID := c.Query("kelas_id")
	tanggal := c.Query("tanggal")

	if kelasID == "" || tanggal == "" {
		c.JSON(http.StatusBadRequest, dto.Error("Parameter kelas_id dan tanggal wajib"))
		return
	}

	// Get all siswa in the class
	var siswa []model.Siswa
	database.DB.Where("kelas_id = ?", kelasID).Find(&siswa)

	// Get absensi for those siswa on that date
	var absensiList []model.Absensi
	database.DB.Preload("Siswa").
		Where("tanggal = ?", tanggal).
		Where("siswa_id IN (SELECT id FROM siswa WHERE kelas_id = ?)", kelasID).
		Find(&absensiList)

	// Map existing absensi by siswa_id
	absensiMap := make(map[uint]model.Absensi)
	for _, a := range absensiList {
		absensiMap[a.SiswaID] = a
	}

	// Build result: combine siswa with their absensi
	type AbsensiWithSiswa struct {
		model.Absensi
		Siswa *model.Siswa `json:"siswa,omitempty"`
	}

	var result []AbsensiWithSiswa
	for _, s := range siswa {
		if a, exists := absensiMap[s.ID]; exists {
			result = append(result, AbsensiWithSiswa{Absensi: a, Siswa: &s})
		} else {
			// No absensi record yet for this student
			result = append(result, AbsensiWithSiswa{
				Absensi: model.Absensi{
					SiswaID: s.ID,
					Tanggal: tanggal,
					Status:  "",
				},
				Siswa: &s,
			})
		}
	}

	c.JSON(http.StatusOK, dto.OK(result))
}

func CreateAbsensiBatch(c *gin.Context) {
	var req dto.CreateAbsensiBatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Error("Input tidak valid: "+err.Error()))
		return
	}

	// Delete existing absensi for these siswa on this date
	for _, item := range req.Data {
		database.DB.Where("siswa_id = ? AND tanggal = ?", item.SiswaID, req.Tanggal).
			Delete(&model.Absensi{})
	}

	// Create new records
	for _, item := range req.Data {
		absensi := model.Absensi{
			SiswaID:    item.SiswaID,
			Tanggal:    req.Tanggal,
			Status:     item.Status,
			Keterangan: item.Keterangan,
		}
		database.DB.Create(&absensi)
	}

	// Return updated data
	kelasID := fmt.Sprintf("%d", req.KelasID)
	tanggal := req.Tanggal
	c.Request.URL.RawQuery = "kelas_id=" + kelasID + "&tanggal=" + tanggal
	GetAbsensi(c)
}

func RekapAbsensiBulanan(c *gin.Context) {
	kelasID := c.Query("kelas_id")
	bulan := c.Query("bulan")
	tahun := c.Query("tahun")

	if kelasID == "" || bulan == "" || tahun == "" {
		c.JSON(http.StatusBadRequest, dto.Error("Parameter kelas_id, bulan, dan tahun wajib"))
		return
	}

	// Build date range
	startDate := fmt.Sprintf("%s-%s-01", tahun, fmt.Sprintf("%02s", bulan))
	endDate := fmt.Sprintf("%s-%s-31", tahun, fmt.Sprintf("%02s", bulan))

	var result []dto.AbsensiRekapBulanan

	rows, err := database.DB.Table("absensi").
		Select(`siswa.id as siswa_id, siswa.nama as nama_siswa, siswa.nis,
			COUNT(CASE WHEN absensi.status = 'hadir' THEN 1 END) as hadir,
			COUNT(CASE WHEN absensi.status = 'sakit' THEN 1 END) as sakit,
			COUNT(CASE WHEN absensi.status = 'izin' THEN 1 END) as izin,
			COUNT(CASE WHEN absensi.status = 'alfa' THEN 1 END) as alfa,
			COUNT(absensi.id) as total`).
		Joins("JOIN siswa ON siswa.id = absensi.siswa_id").
		Where("siswa.kelas_id = ?", kelasID).
		Where("absensi.tanggal >= ? AND absensi.tanggal <= ?", startDate, endDate).
		Group("siswa.id, siswa.nama, siswa.nis").
		Order("siswa.nama asc").
		Rows()

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error("Gagal mengambil data"))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r dto.AbsensiRekapBulanan
		rows.Scan(&r.SiswaID, &r.NamaSiswa, &r.NIS, &r.Hadir, &r.Sakit, &r.Izin, &r.Alfa, &r.Total)
		result = append(result, r)
	}

	c.JSON(http.StatusOK, dto.OK(result))
}

func RekapAbsensiHarian(c *gin.Context) {
	tanggal := c.Query("tanggal")
	if tanggal == "" {
		tanggal = time.Now().Format("2006-01-02")
	}

	type HarianStats struct {
		Status string `json:"status"`
		Total  int64  `json:"total"`
	}

	var stats []HarianStats
	database.DB.Model(&model.Absensi{}).
		Select("status, COUNT(*) as total").
		Where("tanggal = ?", tanggal).
		Group("status").
		Find(&stats)

	result := map[string]int64{
		"hadir": 0, "sakit": 0, "izin": 0, "alfa": 0,
	}
	for _, s := range stats {
		result[s.Status] = s.Total
	}

	c.JSON(http.StatusOK, dto.OK(result))
}
