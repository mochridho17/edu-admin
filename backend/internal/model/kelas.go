package model

import "time"

type Kelas struct {
	ID            uint         `gorm:"primaryKey" json:"id"`
	Nama          string       `gorm:"size:50;not null" json:"nama"`
	Tingkat       string       `gorm:"size:5;not null" json:"tingkat"`
	TahunAjaranID uint         `gorm:"not null" json:"tahun_ajaran_id"`
	WaliKelas     string       `gorm:"size:100" json:"wali_kelas"`
	TahunAjaran   *TahunAjaran `gorm:"foreignKey:TahunAjaranID" json:"tahun_ajaran,omitempty"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
}

func (Kelas) TableName() string {
	return "kelas"
}
