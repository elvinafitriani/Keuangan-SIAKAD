package repository

import (
	"keuangan/entity"
	"strconv"
	"time"

	"gorm.io/gorm"
)

func NewRepoBln(db *gorm.DB) Db {
	return Db{
		DbBln: db,
	}
}

type Db struct {
	DbBln *gorm.DB
}

func (db Db) CreateLapBulanan(lapb entity.LapBulanan) error {
	bulanint, errconv := strconv.Atoi(lapb.Bulan)
	if errconv != nil {
		return errconv
	}
	bulan := time.Month(bulanint)

	tahun, errconv := strconv.Atoi(lapb.Tahun)
	if errconv != nil {
		return errconv
	}

	var dftr []entity.Pendaftaran
	var pemb []entity.Pembangunan
	var smstr []entity.Semesteran
	var sidang []entity.Sidang
	var wisuda []entity.Wisuda
	var total, fieldTotal int

	if err := db.DbBln.First(&dftr, "EXTRACT(MONTH FROM pendaftarans.created_at) = ? AND EXTRACT(YEAR FROM pendaftarans.created_at) = ?", bulan, tahun).Error; err != nil {
		if err := db.DbBln.First(&pemb, "EXTRACT(MONTH FROM pembangunans.created_at) = ? AND EXTRACT(YEAR FROM pembangunans.created_at) = ?", bulan, tahun).Error; err != nil {
			if err := db.DbBln.First(&smstr, "EXTRACT(MONTH FROM semesterans.created_at) = ? AND EXTRACT(YEAR FROM semesterans.created_at) = ?", bulan, tahun).Error; err != nil {
				if err := db.DbBln.First(&sidang, "EXTRACT(MONTH FROM sidangs.created_at) = ? AND EXTRACT(YEAR FROM sidangs.created_at) = ?", bulan, tahun).Error; err != nil {
					if err := db.DbBln.First(&wisuda, "EXTRACT(MONTH FROM wisudas.created_at) = ? AND EXTRACT(YEAR FROM wisudas.created_at) = ?", bulan, tahun).Error; err != nil {
						return err
					}
				}
			}
		}
	}

	if err := db.DbBln.Where("EXTRACT(MONTH FROM pendaftarans.created_at) = ? AND EXTRACT(YEAR FROM pendaftarans.created_at) = ?", bulan, tahun).Find(&dftr).Error; err != nil {
		return err
	}
	if err := db.DbBln.Where("EXTRACT(MONTH FROM pembangunans.created_at) = ? AND EXTRACT(YEAR FROM pembangunans.created_at) = ?", bulan, tahun).Find(&pemb).Error; err != nil {
		return err
	}
	if err := db.DbBln.Where("EXTRACT(MONTH FROM semesterans.created_at) = ? AND EXTRACT(YEAR FROM semesterans.created_at) = ?", bulan, tahun).Find(&smstr).Error; err != nil {
		return err
	}
	if err := db.DbBln.Where("EXTRACT(MONTH FROM sidangs.created_at) = ? AND EXTRACT(YEAR FROM sidangs.created_at) = ?", bulan, tahun).Find(&sidang).Error; err != nil {
		return err
	}
	if err := db.DbBln.Where("EXTRACT(MONTH FROM wisudas.created_at) = ? AND EXTRACT(YEAR FROM wisudas.created_at) = ?", bulan, tahun).Find(&wisuda).Error; err != nil {
		return err
	}

	for _, v := range dftr {
		total += v.Jumlah
	}
	fieldTotal += total
	lapb.Pendaftaran = total

	total = 0
	for _, v := range pemb {
		total += v.Jumlah
	}
	fieldTotal += total
	lapb.Pembangunan = total

	total = 0
	for _, v := range smstr {
		total += v.Jumlah
	}
	fieldTotal += total
	lapb.Semesteran = total

	total = 0
	for _, v := range sidang {
		total += v.Jumlah
	}
	fieldTotal += total
	lapb.Sidang = total

	total = 0
	for _, v := range wisuda {
		total += v.Jumlah
	}
	fieldTotal += total
	lapb.Wisuda = total

	lapb.Total = fieldTotal

	err := db.DbBln.Create(&lapb).Error

	if lapb.Pendaftaran == 0 {
		return err
	}

	if err != nil {
		return err
	}

	return nil

}

func (db Db) GetAllLapBulanan() (lapb []entity.LapBulanan, err error) {
	err = db.DbBln.Find(&lapb).Error

	if err != nil {
		return nil, err
	}

	return lapb, nil
}

func (db Db) GetLapBulananByBulan(bulan string, tahun string) ([]entity.LapBulanan, error) {
	var lapb []entity.LapBulanan
	if err := db.DbBln.First(&lapb, "bulan = ? AND tahun = ?", bulan, tahun).Error; err != nil {
		return nil, err
	}
	err := db.DbBln.Find(&lapb, "bulan = ? AND tahun = ?", bulan, tahun).Error

	if err != nil {
		return nil, err
	}

	return lapb, nil
}

func (db Db) DeleteLapBulanan(bulan string, tahun string) error {
	if err := db.DbBln.First(&entity.LapBulanan{}, "bulan=? AND tahun = ?", bulan, tahun).Error; err != nil {
		return err
	}

	if err := db.DbBln.Where("bulan=? AND tahun = ?", bulan, tahun).Delete(&entity.LapBulanan{}).Error; err != nil {
		return err
	}

	return nil
}
