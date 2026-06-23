package model

import "time"

type Nilai struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	SiswaID       uint           `gorm:"not null;uniqueIndex:idx_nilai_unique" json:"siswa_id"`
	MapelID       uint           `gorm:"not null;uniqueIndex:idx_nilai_unique" json:"mapel_id"`
	Semester      string         `gorm:"size:10;not null;uniqueIndex:idx_nilai_unique" json:"semester"`
	TahunAjaranID uint           `gorm:"not null;uniqueIndex:idx_nilai_unique" json:"tahun_ajaran_id"`
	Tugas         *float64       `gorm:"type:decimal(5,2)" json:"tugas"`
	UTS           *float64       `gorm:"type:decimal(5,2)" json:"uts"`
	UAS           *float64       `gorm:"type:decimal(5,2)" json:"uas"`
	Akhir         *float64       `gorm:"type:decimal(5,2)" json:"akhir"`
	Siswa         *Siswa         `gorm:"foreignKey:SiswaID" json:"siswa,omitempty"`
	Mapel         *MataPelajaran `gorm:"foreignKey:MapelID" json:"mapel,omitempty"`
	TahunAjaran   *TahunAjaran   `gorm:"foreignKey:TahunAjaranID" json:"tahun_ajaran,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

func (Nilai) TableName() string {
	return "nilai"
}
