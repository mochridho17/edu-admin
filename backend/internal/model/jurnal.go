package model

import "time"

type JurnalGuru struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Tanggal       string         `gorm:"size:10;not null" json:"tanggal"`
	NamaGuru      string         `gorm:"size:150;not null" json:"nama_guru"`
	MapelID       uint           `gorm:"not null" json:"mapel_id"`
	KelasID       uint           `gorm:"not null" json:"kelas_id"`
	Kegiatan      string         `gorm:"type:text;not null" json:"kegiatan"`
	Semester      string         `gorm:"size:10;not null" json:"semester"`
	TahunAjaranID uint           `gorm:"not null" json:"tahun_ajaran_id"`
	Mapel         *MataPelajaran `gorm:"foreignKey:MapelID" json:"mapel,omitempty"`
	Kelas         *Kelas         `gorm:"foreignKey:KelasID" json:"kelas,omitempty"`
	TahunAjaran   *TahunAjaran   `gorm:"foreignKey:TahunAjaranID" json:"tahun_ajaran,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

func (JurnalGuru) TableName() string {
	return "jurnal_guru"
}
