package repository

import (
	"errors"
	"fmt"
	"keuangan/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRepoPemb(db *gorm.DB) Db {
	return Db{
		DbPemb: db,
	}
}

type Db struct {
	DbPemb *gorm.DB
}

func (db Db) CreatePembangunan(pemb entity.Pembangunan, ctx *gin.Context) error {
	var client = &http.Client{}
	url := fmt.Sprintf("http://localhost:8081/api/mahasiswa/%s", pemb.Nim)

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

	if err := db.DbPemb.Create(&pemb).Error; err != nil {
		return err
	}
	return nil
}

func (db Db) GetAllPembangunan() (pemb []entity.Pembangunan, err error) {
	err = db.DbPemb.Find(&pemb).Error

	if err != nil {
		return nil, err
	}

	return pemb, nil
}

func (db Db) GetPembangunanByNim(nim string) ([]entity.Pembangunan, error) {
	var pembangunan []entity.Pembangunan
	if err := db.DbPemb.First(&pembangunan, "nim=?", nim).Error; err != nil {
		return nil, err
	}
	err := db.DbPemb.Where("nim = ?", nim).Find(&pembangunan).Error

	if err != nil {
		return nil, err
	}

	return pembangunan, nil
}

func (db Db) UpdatePembangunan(Pemb entity.Pembangunan, nim string, ctx *gin.Context) error {
	if err := db.DbPemb.First(&entity.Pembangunan{}, "nim=?", nim).Error; err != nil {
		return err
	}

	var client = &http.Client{}
	token := ctx.Request.Header.Get("Authorization")

	if Pemb.Nim != "" {
		url := fmt.Sprintf("http://localhost:8081/api/mahasiswa/%s", Pemb.Nim)

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
	if err := db.DbPemb.Where("nim = ?", nim).Updates(&Pemb).Error; err != nil {
		return err
	}
	return nil
}

func (db Db) DeletePembangunan(nim string) error {
	if err := db.DbPemb.First(&entity.Pembangunan{}, "nim=?", nim).Error; err != nil {
		return err
	}

	if err := db.DbPemb.Where("nim=?", nim).Delete(&entity.Pembangunan{}).Error; err != nil {
		return err
	}

	return nil
}
