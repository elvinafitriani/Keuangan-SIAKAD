package repository

import (
	"errors"
	"fmt"
	"keuangan/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRepoSdg(db *gorm.DB) Db {
	return Db{
		DbSdg: db,
	}
}

type Db struct {
	DbSdg *gorm.DB
}

func (db Db) CreateSidang(sdg entity.Sidang, ctx *gin.Context) error {
	var client = &http.Client{}
	url := fmt.Sprintf("http://localhost:8081/api/mahasiswa/%s", sdg.Nim)
	token := ctx.Request.Header.Get("Authorization")

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

	if err := db.DbSdg.Create(&sdg).Error; err != nil {
		return err
	}
	return nil
}

func (db Db) GetAllSidang() (sdg []entity.Sidang, err error) {
	err = db.DbSdg.Find(&sdg).Error

	if err != nil {
		return nil, err
	}

	return sdg, nil
}

func (db Db) GetSidangByNim(nim string) ([]entity.Sidang, error) {
	var sidang []entity.Sidang
	if err := db.DbSdg.First(&sidang, "nim=?", nim).Error; err != nil {
		return nil, err
	}
	err := db.DbSdg.Find(&sidang, "nim = ?", nim).Error

	if err != nil {
		return nil, err
	}

	return sidang, nil
}

func (db Db) UpdateSidang(Sdg entity.Sidang, nim string, ctx *gin.Context) error {
	if err := db.DbSdg.First(&entity.Sidang{}, "nim=?", nim).Error; err != nil {
		return err
	}

	token := ctx.Request.Header.Get("Authorization")

	var client = &http.Client{}
	if Sdg.Nim != "" {
		url := fmt.Sprintf("http://localhost:8081/api/mahasiswa/%s", Sdg.Nim)

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

	if err := db.DbSdg.Where("nim=?", nim).Updates(&Sdg).Error; err != nil {
		return err
	}
	return nil
}

func (db Db) DeleteSidang(nim string) error {
	if err := db.DbSdg.First(&entity.Sidang{}, "nim=?", nim).Error; err != nil {
		return err
	}

	if err := db.DbSdg.Where("nim=?", nim).Delete(&entity.Sidang{}).Error; err != nil {
		return err
	}

	return nil
}
