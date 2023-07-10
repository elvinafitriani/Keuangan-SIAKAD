package repository

import (
	"errors"
	"fmt"
	"keuangan/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRepoDftr(db *gorm.DB) Db {
	return Db{
		DbDftr: db,
	}
}

type Db struct {
	DbDftr *gorm.DB
}

func (db Db) CreatePendaftaran(dftr entity.Pendaftaran, ctx *gin.Context) error {
	var client = &http.Client{}
	url := fmt.Sprintf("http://localhost:8081/api/mahasiswa/%s", dftr.Nim)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	token := ctx.Request.Header.Get("Authorization")

	if err != nil {
		return err
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	res, errRes := client.Do(req)
	if errRes != nil {
		return errRes
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New("bad request")
	}

	if err := db.DbDftr.Create(&dftr).Error; err != nil {
		return err
	}
	return nil
}

func (db Db) GetAllPendaftaran() (dftr []entity.Pendaftaran, err error) {
	err = db.DbDftr.Find(&dftr).Error

	if err != nil {
		return nil, err
	}

	return dftr, nil
}

func (db Db) GetPendaftaranByNim(nim string) ([]entity.Pendaftaran, error) {
	var daftar []entity.Pendaftaran
	if err := db.DbDftr.First(&daftar, "nim=?", nim).Error; err != nil {
		return nil, err
	}

	err := db.DbDftr.Find(&daftar, "nim = ?", nim).Error

	if err != nil {
		return nil, err
	}

	return daftar, nil
}

func (db Db) UpdatePendaftaran(Dftr entity.Pendaftaran, nim string, ctx *gin.Context) error {
	if err := db.DbDftr.First(&entity.Pendaftaran{}, "nim=?", nim).Error; err != nil {
		return err
	}

	var client = &http.Client{}

	token := ctx.Request.Header.Get("Authorization")

	if Dftr.Nim != "" {
		url := fmt.Sprintf("http://localhost:8081/api/mahasiswa/%s", Dftr.Nim)

		req, err := http.NewRequest(http.MethodGet, url, nil)

		if err != nil {
			return err
		}

		req.Header.Set("Authorization", token)
		req.Header.Set("Content-Type", "application/json")

		res, errRes := client.Do(req)
		if errRes != nil {
			return errRes
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return errors.New("bad request")
		}
	}

	if err := db.DbDftr.Where("nim=?", nim).Updates(&Dftr).Error; err != nil {
		return err
	}
	return nil
}

func (db Db) DeletePendaftaran(nim string) error {
	if err := db.DbDftr.First(&entity.Pendaftaran{}, "nim=?", nim).Error; err != nil {
		return err
	}

	if err := db.DbDftr.Where("nim=?", nim).Delete(&entity.Pendaftaran{}).Error; err != nil {
		return err
	}

	return nil
}
