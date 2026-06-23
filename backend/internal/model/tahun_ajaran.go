package model

import "time"

type TahunAjaran struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `gorm:"size:20;not null" json:"nama"`
	Semester  string    `gorm:"size:10;not null" json:"semester"`
	IsActive  bool      `gorm:"default:false" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (TahunAjaran) TableName() string {
	return "tahun_ajaran"
}
