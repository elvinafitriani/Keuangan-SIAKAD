package entity

import "gorm.io/gorm"

type Pendaftaran struct {
	gorm.Model
	Nim    string `gorm:"type:char(11)" json:"nim" binding:"min=11"`
	Jumlah int    `json:"jumlah" binding:"required"`
}

type Pembangunan struct {
	gorm.Model
	Nim    string `gorm:"type:char(11)" json:"nim" binding:"min=11"`
	Jumlah int    `json:"jumlah" binding:"required"`
}

type Semesteran struct {
	gorm.Model
	Nim    string `gorm:"type:char(11)" json:"nim" binding:"min=11"`
	Jumlah int    `json:"jumlah" binding:"required"`
}

type Sidang struct {
	gorm.Model
	Nim    string `gorm:"type:char(11)" json:"nim" binding:"min=11"`
	Jumlah int    `json:"jumlah" binding:"required"`
}

type Wisuda struct {
	gorm.Model
	Nim    string `gorm:"type:char(11)" json:"nim" binding:"min=11"`
	Jumlah int    `json:"jumlah" binding:"required"`
}

type LapBulanan struct {
	gorm.Model
	Bulan       string `gorm:"type:char(2)" json:"bulan" binding:"required"`
	Tahun       string `gorm:"type:char(4)" json:"tahun" binding:"required,min=4"`
	Pendaftaran int    `json:"pendaftaran"`
	Pembangunan int    `json:"pembangunan"`
	Semesteran  int    `json:"semesteran"`
	Sidang      int    `json:"sidang"`
	Wisuda      int    `json:"wisuda"`
	Total       int    `json:"total"`
}

type LapTahunan struct {
	gorm.Model
	Tahun       string `gorm:"type:char(4)" json:"tahun" binding:"required,min=4"`
	Pendaftaran int    `json:"pendaftaran"`
	Pembangunan int    `json:"pembangunan"`
	Semesteran  int    `json:"semesteran"`
	Sidang      int    `json:"sidang"`
	Wisuda      int    `json:"wisuda"`
	Total       int    `json:"total"`
}
