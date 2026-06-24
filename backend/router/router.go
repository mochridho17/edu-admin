package router

import (
	"edu-admin/internal/config"
	"edu-admin/internal/handler"
	"edu-admin/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Global middleware
	r.Use(middleware.CORSMiddleware())

	// Public routes
	api := r.Group("/api")
	{
		api.POST("/auth/login", handler.Login)
		api.POST("/auth/refresh", handler.Login)
	}

	// Protected routes
	auth := api.Group("")
	auth.Use(middleware.AuthMiddleware(config.AppConfig.JWTSecret))
	{
		// Dashboard
		auth.GET("/dashboard/stats", handler.GetDashboardStats)

		// Kelas
		auth.GET("/kelas", handler.GetKelas)
		auth.POST("/kelas", handler.CreateKelas)
		auth.PUT("/kelas/:id", handler.UpdateKelas)
		auth.DELETE("/kelas/:id", handler.DeleteKelas)

		// Siswa
		auth.GET("/siswa", handler.GetSiswa)
		auth.POST("/siswa", handler.CreateSiswa)
		auth.PUT("/siswa/:id", handler.UpdateSiswa)
		auth.DELETE("/siswa/:id", handler.DeleteSiswa)

		// Mata Pelajaran
		auth.GET("/mapel", handler.GetMapel)
		auth.POST("/mapel", handler.CreateMapel)
		auth.PUT("/mapel/:id", handler.UpdateMapel)
		auth.DELETE("/mapel/:id", handler.DeleteMapel)

		// Absensi
		auth.GET("/absensi", handler.GetAbsensi)
		auth.POST("/absensi", handler.CreateAbsensiBatch)
		auth.GET("/absensi/rekap/bulanan", handler.RekapAbsensiBulanan)
		auth.GET("/absensi/rekap/harian", handler.RekapAbsensiHarian)

		// Nilai
		auth.GET("/nilai", handler.GetNilai)
		auth.POST("/nilai/batch", handler.CreateNilaiBatch)
		auth.GET("/nilai/rekap/siswa/:id", handler.RekapNilaiSiswa)
		auth.GET("/nilai/rekap/kelas", handler.RankingKelas)

		// Laporan
		// Jurnal Guru
		auth.GET("/jurnal-guru", handler.GetJurnalGuru)
		auth.POST("/jurnal-guru", handler.CreateJurnalGuru)
		auth.PUT("/jurnal-guru/:id", handler.UpdateJurnalGuru)
		auth.DELETE("/jurnal-guru/:id", handler.DeleteJurnalGuru)

		// Laporan
		auth.GET("/laporan/absensi", handler.ExportAbsensi)
		auth.GET("/laporan/nilai", handler.ExportNilai)
		auth.GET("/laporan/jurnal-guru", handler.ExportJurnalGuru)
		// School Info
		auth.GET("/school/info", handler.GetSchoolInfo)

		// Super Admin routes
		super := auth.Group("")
		super.Use(middleware.SuperAdminMiddleware())
		{
			// Users
			super.GET("/users", handler.GetUsers)
			super.POST("/users", handler.CreateUser)
			super.PUT("/users/:id", handler.UpdateUser)
			super.DELETE("/users/:id", handler.DeleteUser)

			// Tahun Ajaran
			super.GET("/tahun-ajaran", handler.GetTahunAjaran)
			super.POST("/tahun-ajaran", handler.CreateTahunAjaran)
			super.PUT("/tahun-ajaran/:id", handler.UpdateTahunAjaran)
			super.DELETE("/tahun-ajaran/:id", handler.DeleteTahunAjaran)
			super.PATCH("/tahun-ajaran/:id/activate", handler.ActivateTahunAjaran)
		}
	}

	return r
}
