package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"edu-admin/internal/database"
	"edu-admin/internal/dto"
	"edu-admin/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func ExportAbsensi(c *gin.Context) {
	kelasID := c.Query("kelas_id")
	bulan := c.Query("bulan")
	tahun := c.Query("tahun")

	if kelasID == "" || bulan == "" || tahun == "" {
		c.JSON(http.StatusBadRequest, dto.Error("Parameter kelas_id, bulan, dan tahun wajib"))
		return
	}

	// Get kelas info
	var kelas model.Kelas
	if err := database.DB.Preload("TahunAjaran").First(&kelas, kelasID).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.Error("Kelas tidak ditemukan"))
		return
	}

	// Get absensi data
	startDate := fmt.Sprintf("%s-%s-01", tahun, fmt.Sprintf("%02s", bulan))
	endDate := fmt.Sprintf("%s-%s-31", tahun, fmt.Sprintf("%02s", bulan))

	type RekapRow struct {
		NIS   string
		Nama  string
		Hadir int64
		Sakit int64
		Izin  int64
		Alfa  int64
		Total int64
	}

	var rows []RekapRow
	database.DB.Table("absensi").
		Select(`siswa.nis, siswa.nama,
			COUNT(CASE WHEN absensi.status = 'hadir' THEN 1 END) as hadir,
			COUNT(CASE WHEN absensi.status = 'sakit' THEN 1 END) as sakit,
			COUNT(CASE WHEN absensi.status = 'izin' THEN 1 END) as izin,
			COUNT(CASE WHEN absensi.status = 'alfa' THEN 1 END) as alfa,
			COUNT(absensi.id) as total`).
		Joins("JOIN siswa ON siswa.id = absensi.siswa_id").
		Where("siswa.kelas_id = ?", kelasID).
		Where("absensi.tanggal >= ? AND absensi.tanggal <= ?", startDate, endDate).
		Group("siswa.nis, siswa.nama").
		Order("siswa.nama asc").
		Scan(&rows)

	// Create Excel
	f := excelize.NewFile()
	defer f.Close()

	// Sheet name
	sheet := "Absensi"
	f.SetSheetName("Sheet1", sheet)

	// Header style
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 12, Color: "#FFFFFF"},
		Fill:      excelize.Fill{Type: "pattern", Pattern: 1, Color: []string{"#009EF7"}},
		Alignment: &excelize.Alignment{Horizontal: "center"},
	})

	// Title
	titleStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 14},
		Alignment: &excelize.Alignment{Horizontal: "center"},
	})
	f.SetCellValue(sheet, "A1", fmt.Sprintf("LAPORAN ABSENSI - %s", kelas.Nama))
	f.SetCellStyle(sheet, "A1", "F1", titleStyle)
	f.MergeCell(sheet, "A1", "F1")

	// Period
	bulanNama := map[int]string{1: "Januari", 2: "Februari", 3: "Maret", 4: "April", 5: "Mei", 6: "Juni",
		7: "Juli", 8: "Agustus", 9: "September", 10: "Oktober", 11: "November", 12: "Desember"}
	b, _ := strconv.Atoi(bulan)
	periode := fmt.Sprintf("Bulan: %s %s", bulanNama[b], tahun)
	f.SetCellValue(sheet, "A2", periode)
	f.SetCellStyle(sheet, "A2", "F2", titleStyle)
	f.MergeCell(sheet, "A2", "F2")

	// Headers
	headers := []string{"No", "NIS", "Nama Siswa", "Hadir", "Sakit", "Izin", "Alfa", "Total"}
	for i, h := range headers {
		col := string(rune('A' + i))
		cell := fmt.Sprintf("%s4", col)
		f.SetCellValue(sheet, cell, h)
		f.SetCellStyle(sheet, cell, cell, headerStyle)
	}

	// Data
	for i, row := range rows {
		r := i + 5
		f.SetCellValue(sheet, fmt.Sprintf("A%d", r), i+1)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", r), row.NIS)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", r), row.Nama)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", r), row.Hadir)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", r), row.Sakit)
		f.SetCellValue(sheet, fmt.Sprintf("F%d", r), row.Izin)
		f.SetCellValue(sheet, fmt.Sprintf("G%d", r), row.Alfa)
		f.SetCellValue(sheet, fmt.Sprintf("H%d", r), row.Total)
	}

	// Column widths
	f.SetColWidth(sheet, "A", "A", 6)
	f.SetColWidth(sheet, "B", "B", 15)
	f.SetColWidth(sheet, "C", "C", 30)
	f.SetColWidth(sheet, "D", "H", 10)

	// Write response
	filename := fmt.Sprintf("Laporan_Absensi_%s_Bulan%d_%s.xlsx", kelas.Nama, b, tahun)
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))

	if err := f.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error("Gagal menulis file Excel"))
	}
}

func ExportJurnalGuru(c *gin.Context) {
	namaGuru := c.Query("nama_guru")
	mapelID := c.Query("mapel_id")
	kelasID := c.Query("kelas_id")
	semester := c.Query("semester")
	taID := c.Query("tahun_ajaran_id")

	query := database.DB.Preload("Mapel").Preload("Kelas").Preload("TahunAjaran").Order("tanggal desc")

	if namaGuru != "" {
		query = query.Where("nama_guru ILIKE ?", "%"+namaGuru+"%")
	}
	if mapelID != "" {
		query = query.Where("mapel_id = ?", mapelID)
	}
	if kelasID != "" {
		query = query.Where("kelas_id = ?", kelasID)
	}
	if semester != "" {
		query = query.Where("semester = ?", semester)
	}
	if taID != "" {
		query = query.Where("tahun_ajaran_id = ?", taID)
	}

	var list []model.JurnalGuru
	query.Find(&list)

	// Create Excel
	f := excelize.NewFile()
	defer f.Close()

	sheet := "Jurnal Guru"
	f.SetSheetName("Sheet1", sheet)

	// Styles
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 12, Color: "#FFFFFF"},
		Fill:      excelize.Fill{Type: "pattern", Pattern: 1, Color: []string{"#009EF7"}},
		Alignment: &excelize.Alignment{Horizontal: "center"},
	})
	titleStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 14},
		Alignment: &excelize.Alignment{Horizontal: "center"},
	})

	// Title
	f.SetCellValue(sheet, "A1", "LAPORAN JURNAL GURU")
	f.SetCellStyle(sheet, "A1", "F1", titleStyle)
	f.MergeCell(sheet, "A1", "F1")

	// Headers
	headers := []string{"No", "Hari/Tanggal", "Nama Guru", "Mata Pelajaran", "Kelas", "Semester", "Kegiatan Pembelajaran"}
	for i, h := range headers {
		col := string(rune('A' + i))
		cell := fmt.Sprintf("%s3", col)
		f.SetCellValue(sheet, cell, h)
		f.SetCellStyle(sheet, cell, cell, headerStyle)
	}

	// Data
	for i, j := range list {
		r := i + 4
		mapelNama := ""
		if j.Mapel != nil {
			mapelNama = j.Mapel.Nama
		}
		kelasNama := ""
		if j.Kelas != nil {
			kelasNama = j.Kelas.Nama
		}
		semesterLabel := "Ganjil"
		if j.Semester == "genap" {
			semesterLabel = "Genap"
		}

		f.SetCellValue(sheet, fmt.Sprintf("A%d", r), i+1)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", r), j.Tanggal)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", r), j.NamaGuru)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", r), mapelNama)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", r), kelasNama)
		f.SetCellValue(sheet, fmt.Sprintf("F%d", r), semesterLabel)
		f.SetCellValue(sheet, fmt.Sprintf("G%d", r), j.Kegiatan)
	}

	// Column widths
	f.SetColWidth(sheet, "A", "A", 6)
	f.SetColWidth(sheet, "B", "B", 16)
	f.SetColWidth(sheet, "C", "C", 25)
	f.SetColWidth(sheet, "D", "D", 25)
	f.SetColWidth(sheet, "E", "E", 15)
	f.SetColWidth(sheet, "F", "F", 12)
	f.SetColWidth(sheet, "G", "G", 50)

	// Write response
	filename := fmt.Sprintf("Laporan_Jurnal_Guru.xlsx")
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))

	if err := f.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error("Gagal menulis file Excel"))
	}
}

func ExportNilai(c *gin.Context) {
	kelasID := c.Query("kelas_id")
	mapelID := c.Query("mapel_id")

	if kelasID == "" || mapelID == "" {
		c.JSON(http.StatusBadRequest, dto.Error("Parameter kelas_id dan mapel_id wajib"))
		return
	}

	// Get kelas & mapel info
	var kelas model.Kelas
	var mapel model.MataPelajaran
	if err := database.DB.First(&kelas, kelasID).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.Error("Kelas tidak ditemukan"))
		return
	}
	if err := database.DB.First(&mapel, mapelID).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.Error("Mapel tidak ditemukan"))
		return
	}

	// Get nilai data
	var nilai []model.Nilai
	database.DB.Preload("Siswa").
		Joins("JOIN siswa ON siswa.id = nilai.siswa_id").
		Where("siswa.kelas_id = ?", kelasID).
		Where("nilai.mapel_id = ?", mapelID).
		Order("siswa.nama asc").
		Find(&nilai)

	// Create Excel
	f := excelize.NewFile()
	defer f.Close()

	sheet := "Nilai"
	f.SetSheetName("Sheet1", sheet)

	// Styles
	headerStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 12, Color: "#FFFFFF"},
		Fill:      excelize.Fill{Type: "pattern", Pattern: 1, Color: []string{"#009EF7"}},
		Alignment: &excelize.Alignment{Horizontal: "center"},
	})
	titleStyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Bold: true, Size: 14},
		Alignment: &excelize.Alignment{Horizontal: "center"},
	})

	// Title
	f.SetCellValue(sheet, "A1", fmt.Sprintf("LAPORAN NILAI - %s", mapel.Nama))
	f.SetCellStyle(sheet, "A1", "G1", titleStyle)
	f.MergeCell(sheet, "A1", "G1")

	f.SetCellValue(sheet, "A2", fmt.Sprintf("Kelas: %s", kelas.Nama))
	f.SetCellStyle(sheet, "A2", "G2", titleStyle)
	f.MergeCell(sheet, "A2", "G2")

	f.SetCellValue(sheet, "A3", fmt.Sprintf("KKM: %d", mapel.KKM))
	f.SetCellStyle(sheet, "A3", "G3", titleStyle)
	f.MergeCell(sheet, "A3", "G3")

	// Headers
	headers := []string{"No", "NIS", "Nama Siswa", "Tugas", "UTS", "UAS", "Akhir"}
	for i, h := range headers {
		col := string(rune('A' + i))
		f.SetCellValue(sheet, fmt.Sprintf("%s5", col), h)
		f.SetCellStyle(sheet, fmt.Sprintf("%s5", col), fmt.Sprintf("%s5", col), headerStyle)
	}

	// Data
	for i, n := range nilai {
		r := i + 6
		tugas := ""
		uts := ""
		uas := ""
		akhir := ""
		if n.Tugas != nil {
			tugas = fmt.Sprintf("%.0f", *n.Tugas)
		}
		if n.UTS != nil {
			uts = fmt.Sprintf("%.0f", *n.UTS)
		}
		if n.UAS != nil {
			uas = fmt.Sprintf("%.0f", *n.UAS)
		}
		if n.Akhir != nil {
			akhir = fmt.Sprintf("%.0f", *n.Akhir)
		}

		f.SetCellValue(sheet, fmt.Sprintf("A%d", r), i+1)
		if n.Siswa != nil {
			f.SetCellValue(sheet, fmt.Sprintf("B%d", r), n.Siswa.NIS)
			f.SetCellValue(sheet, fmt.Sprintf("C%d", r), n.Siswa.Nama)
		}
		f.SetCellValue(sheet, fmt.Sprintf("D%d", r), tugas)
		f.SetCellValue(sheet, fmt.Sprintf("E%d", r), uts)
		f.SetCellValue(sheet, fmt.Sprintf("F%d", r), uas)
		f.SetCellValue(sheet, fmt.Sprintf("G%d", r), akhir)
	}

	// Column widths
	f.SetColWidth(sheet, "A", "A", 6)
	f.SetColWidth(sheet, "B", "B", 15)
	f.SetColWidth(sheet, "C", "C", 30)
	f.SetColWidth(sheet, "D", "G", 10)

	// Write response
	filename := fmt.Sprintf("Laporan_Nilai_%s_%s.xlsx", kelas.Nama, mapel.Nama)
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))

	if err := f.Write(c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error("Gagal menulis file Excel"))
	}
}
