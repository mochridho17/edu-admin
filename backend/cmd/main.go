package main

import (
	"log"

	"edu-admin/internal/config"
	"edu-admin/internal/database"
	"edu-admin/router"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Gagal load config: %v", err)
	}

	// Connect database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Gagal koneksi database: %v", err)
	}

	// Migrate
	if err := database.Migrate(db); err != nil {
		log.Fatalf("Gagal migrasi: %v", err)
	}

	// Seed
	if err := database.Seed(db); err != nil {
		log.Printf("Warning: Seed gagal: %v", err)
	}

	// Setup router
	r := router.SetupRouter()

	// Start server
	addr := ":" + cfg.ServerPort
	log.Printf("🚀 Server berjalan di http://localhost%s", addr)

	if err := r.Run(addr); err != nil {
		log.Fatalf("Gagal start server: %v", err)
	}
}
