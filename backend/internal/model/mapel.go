package model

import "time"

type MataPelajaran struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Kode      string    `gorm:"uniqueIndex;size:30;not null" json:"kode"`
	Nama      string    `gorm:"size:150;not null" json:"nama"`
	KKM       int       `gorm:"default:75" json:"kkm"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (MataPelajaran) TableName() string {
	return "mata_pelajaran"
}
