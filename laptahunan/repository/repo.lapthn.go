package repository

import (
	"keuangan/entity"
	"strconv"

	"gorm.io/gorm"
)

func NewRepoThn(db *gorm.DB) Db {
	return Db{
		DbThn: db,
	}
}

type Db struct {
	DbThn *gorm.DB
}

func (db Db) CreateLapTahunan(lapt entity.LapTahunan) error {
	tahun, errconv := strconv.Atoi(lapt.Tahun)
	if errconv != nil {
		return errconv
	}

	var dftr []entity.Pendaftaran
	var pemb []entity.Pembangunan
	var smstr []entity.Semesteran
	var sidang []entity.Sidang
	var wisuda []entity.Wisuda
	var total, fieldTotal int

	if err := db.DbThn.First(&dftr, "EXTRACT(YEAR FROM pendaftarans.created_at) = ?", tahun).Error; err != nil {
		if err := db.DbThn.First(&pemb, "EXTRACT(YEAR FROM pembangunans.created_at) = ?", tahun).Error; err != nil {
			if err := db.DbThn.First(&smstr, "EXTRACT(YEAR FROM semesterans.created_at) = ?", tahun).Error; err != nil {
				if err := db.DbThn.First(&sidang, "EXTRACT(YEAR FROM sidangs.created_at) = ?", tahun).Error; err != nil {
					if err := db.DbThn.First(&wisuda, "EXTRACT(YEAR FROM wisudas.created_at) = ?", tahun).Error; err != nil {
						return err
					}
				}
			}
		}
	}

	if err := db.DbThn.Where("EXTRACT(YEAR FROM pendaftarans.created_at) = ?", tahun).Find(&dftr).Error; err != nil {
		return err
	}
	if err := db.DbThn.Where("EXTRACT(YEAR FROM pembangunans.created_at) = ?", tahun).Find(&pemb).Error; err != nil {
		return err
	}
	if err := db.DbThn.Where("EXTRACT(YEAR FROM semesterans.created_at) = ?", tahun).Find(&smstr).Error; err != nil {
		return err
	}
	if err := db.DbThn.Where("EXTRACT(YEAR FROM sidangs.created_at) = ?", tahun).Find(&sidang).Error; err != nil {
		return err
	}
	if err := db.DbThn.Where("EXTRACT(YEAR FROM wisudas.created_at) = ?", tahun).Find(&wisuda).Error; err != nil {
		return err
	}

	for _, v := range dftr {
		total += v.Jumlah
	}
	fieldTotal += total
	lapt.Pendaftaran = total

	total = 0
	for _, v := range pemb {
		total += v.Jumlah
	}
	fieldTotal += total
	lapt.Pembangunan = total

	total = 0
	for _, v := range smstr {
		total += v.Jumlah
	}
	fieldTotal += total
	lapt.Semesteran = total

	total = 0
	for _, v := range sidang {
		total += v.Jumlah
	}
	fieldTotal += total
	lapt.Sidang = total

	total = 0
	for _, v := range wisuda {
		total += v.Jumlah
	}
	fieldTotal += total
	lapt.Wisuda = total

	lapt.Total = fieldTotal

	err := db.DbThn.Create(&lapt).Error

	if err != nil {
		return err
	}

	return nil

}

func (db Db) GetAllLapTahunan() (lapt []entity.LapTahunan, err error) {
	err = db.DbThn.Find(&lapt).Error

	if err != nil {
		return nil, err
	}

	return lapt, nil
}

func (db Db) GetLapTahunanByTahun(tahun string) ([]entity.LapTahunan, error) {
	var lapt []entity.LapTahunan
	if err := db.DbThn.First(&lapt, "tahun = ?", tahun).Error; err != nil {
		return nil, err
	}

	err := db.DbThn.Find(&lapt, "tahun = ?", tahun).Error

	if err != nil {
		return nil, err
	}

	return lapt, nil
}

func (db Db) DeleteLapTahunan(tahun string) error {
	if err := db.DbThn.First(&entity.LapTahunan{}, "tahun=?", tahun).Error; err != nil {
		return err
	}

	if err := db.DbThn.Where("tahun=?", tahun).Delete(&entity.LapTahunan{}).Error; err != nil {
		return err
	}

	return nil
}
