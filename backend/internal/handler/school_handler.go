package handler

import (
	"net/http"
	"os"

	"edu-admin/internal/dto"

	"github.com/gin-gonic/gin"
)

func GetSchoolInfo(c *gin.Context) {
	schoolName := os.Getenv("SCHOOL_NAME")
	if schoolName == "" {
		schoolName = "SDN 01 Menteng"
	}
	npsn := os.Getenv("SCHOOL_NPSN")
	if npsn == "" {
		npsn = "20105678"
	}
	headmaster := os.Getenv("SCHOOL_HEADMASTER")
	if headmaster == "" {
		headmaster = "Hj. Siti Aminah, S.Pd."
	}
	teachers := os.Getenv("SCHOOL_TEACHERS")
	if teachers == "" {
		teachers = "12 Orang"
	}
	accreditation := os.Getenv("SCHOOL_ACCEREDITATION")
	if accreditation == "" {
		accreditation = "A (Unggul)"
	}

	c.JSON(http.StatusOK, dto.OK(map[string]string{
		"nama_sekolah":   schoolName,
		"npsn":           npsn,
		"kepala_sekolah": headmaster,
		"jumlah_guru":    teachers,
		"akreditasi":     accreditation,
	}))
}
