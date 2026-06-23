package database

import (
	"database/sql"
	"fmt"
	"log"

	"edu-admin/internal/config"
	"edu-admin/internal/model"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(cfg *config.Config) (*gorm.DB, error) {
	// First, try to create database if not exists
	createDatabaseIfNotExists(cfg)

	// Now connect to the actual database
	var err error
	DB, err = gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		return nil, fmt.Errorf("gagal konek database: %w", err)
	}

	sqlDB, err := DB.DB()
	if err == nil {
		sqlDB.SetMaxOpenConns(10)
		sqlDB.SetMaxIdleConns(2)
	}

	log.Println("✅ Database terhubung")
	return DB, nil
}

func createDatabaseIfNotExists(cfg *config.Config) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBSSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("Warning: tidak bisa konek ke postgres default: %v", err)
		return
	}
	defer db.Close()

	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", cfg.DBName).Scan(&exists)
	if err != nil {
		log.Printf("Warning: gagal cek database: %v", err)
		return
	}

	if !exists {
		_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", cfg.DBName))
		if err != nil {
			log.Printf("Warning: gagal buat database: %v", err)
			return
		}
		log.Printf("✅ Database '%s' berhasil dibuat", cfg.DBName)
	}
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.User{},
		&model.TahunAjaran{},
		&model.Kelas{},
		&model.Siswa{},
		&model.MataPelajaran{},
		&model.Absensi{},
		&model.Nilai{},
	)
	if err != nil {
		return fmt.Errorf("gagal migrasi: %w", err)
	}
	log.Println("✅ Migrasi database selesai")
	return nil
}

func Seed(db *gorm.DB) error {
	var count int64
	db.Model(&model.User{}).Count(&count)
	if count > 0 {
		log.Println("ℹ️  Seed sudah ada, skip")
		return nil
	}

	hash, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	admin := model.User{
		Username: "admin",
		Email:    "admin@sekolah.sch.id",
		Password: string(hash),
		Role:     "super_admin",
	}
	if err := db.Create(&admin).Error; err != nil {
		return err
	}

	ta := model.TahunAjaran{
		Nama:     "2025/2026",
		Semester: "ganjil",
		IsActive: true,
	}
	if err := db.Create(&ta).Error; err != nil {
		return err
	}

	log.Println("✅ Seed data berhasil")
	log.Println("   User: admin / admin123 (Super Admin)")
	return nil
}
