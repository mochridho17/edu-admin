package model

import "time"

type Absensi struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	SiswaID    uint      `gorm:"not null;uniqueIndex:idx_absensi_siswa_tanggal" json:"siswa_id"`
	Tanggal    string    `gorm:"size:10;not null;uniqueIndex:idx_absensi_siswa_tanggal" json:"tanggal"`
	Status     string    `gorm:"size:10;not null" json:"status"`
	Keterangan string    `gorm:"size:255" json:"keterangan"`
	Siswa      *Siswa    `gorm:"foreignKey:SiswaID" json:"siswa,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (Absensi) TableName() string {
	return "absensi"
}
