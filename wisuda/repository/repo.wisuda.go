package repository

import (
	"errors"
	"fmt"
	"keuangan/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRepoWsd(db *gorm.DB) Db {
	return Db{
		DbWsd: db,
	}
}

type Db struct {
	DbWsd *gorm.DB
}

func (db Db) CreateWisuda(wsd entity.Wisuda, ctx *gin.Context) error {
	var client = &http.Client{}
	url := fmt.Sprintf("http://localhost:8081/api/mahasiswa/%s", wsd.Nim)
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

	if err := db.DbWsd.Create(&wsd).Error; err != nil {
		return err
	}
	return nil
}

func (db Db) GetAllWisuda() (wsd []entity.Wisuda, err error) {
	err = db.DbWsd.Find(&wsd).Error

	if err != nil {
		return nil, err
	}

	return wsd, nil
}

func (db Db) GetWisudaByNim(nim string) ([]entity.Wisuda, error) {
	var wisuda []entity.Wisuda
	if err := db.DbWsd.First(&wisuda, "nim=?", nim).Error; err != nil {
		return nil, err
	}
	err := db.DbWsd.Find(&wisuda, "nim = ?", nim).Error

	if err != nil {
		return nil, err
	}

	return wisuda, nil
}

func (db Db) UpdateWisuda(Wsd entity.Wisuda, nim string, ctx *gin.Context) error {
	if err := db.DbWsd.First(&entity.Wisuda{}, "nim=?", nim).Error; err != nil {
		return err
	}
	token := ctx.Request.Header.Get("Authorization")

	var client = &http.Client{}
	if Wsd.Nim != "" {
		url := fmt.Sprintf("http://localhost:8081/api/mahasiswa/%s", Wsd.Nim)

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

	if err := db.DbWsd.Where("nim=?", nim).Updates(&Wsd).Error; err != nil {
		return err
	}
	return nil
}

func (db Db) DeleteWisuda(nim string) error {
	if err := db.DbWsd.First(&entity.Wisuda{}, "nim=?", nim).Error; err != nil {
		return err
	}

	if err := db.DbWsd.Where("nim=?", nim).Delete(&entity.Wisuda{}).Error; err != nil {
		return err
	}

	return nil
}
