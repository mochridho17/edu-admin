package model

import "time"

type Siswa struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	NIS          string    `gorm:"uniqueIndex;size:30;not null" json:"nis"`
	Nama         string    `gorm:"size:150;not null" json:"nama"`
	JenisKelamin string    `gorm:"size:2;not null" json:"jenis_kelamin"`
	KelasID      uint      `gorm:"not null" json:"kelas_id"`
	Kelas        *Kelas    `gorm:"foreignKey:KelasID" json:"kelas,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (Siswa) TableName() string {
	return "siswa"
}
