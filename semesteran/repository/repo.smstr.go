package repository

import (
	"errors"
	"fmt"
	"keuangan/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRepoSmstr(db *gorm.DB) Db {
	return Db{
		DbSmstr: db,
	}
}

type Db struct {
	DbSmstr *gorm.DB
}

func (db Db) CreateSemesteran(smstr entity.Semesteran, ctx *gin.Context) error {
	var client = &http.Client{}
	url := fmt.Sprintf("http://localhost:8081/api/mahasiswa/%s", smstr.Nim)
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

	if err := db.DbSmstr.Create(&smstr).Error; err != nil {
		return err
	}
	return nil
}

func (db Db) GetAllSemesteran() (smstr []entity.Semesteran, err error) {
	err = db.DbSmstr.Find(&smstr).Error

	if err != nil {
		return nil, err
	}

	return smstr, nil
}

func (db Db) GetSemesteranByNim(nim string) ([]entity.Semesteran, error) {
	var semesteran []entity.Semesteran
	if err := db.DbSmstr.First(&semesteran, "nim=?", nim).Error; err != nil {
		return nil, err
	}

	err := db.DbSmstr.Find(&semesteran, "nim = ?", nim).Error

	if err != nil {
		return nil, err
	}

	return semesteran, nil
}

func (db Db) UpdateSemesteran(Smstr entity.Semesteran, nim string, ctx *gin.Context) error {
	if err := db.DbSmstr.First(&entity.Semesteran{}, "nim=?", nim).Error; err != nil {
		return err
	}

	var client = &http.Client{}
	token := ctx.Request.Header.Get("Authorization")

	if Smstr.Nim != "" {
		url := fmt.Sprintf("http://localhost:8081/api/mahasiswa/%s", Smstr.Nim)

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
	if err := db.DbSmstr.Where("nim=?", nim).Updates(&Smstr).Error; err != nil {
		return err
	}
	return nil
}

func (db Db) DeleteSemesteran(nim string) error {
	if err := db.DbSmstr.First(&entity.Semesteran{}, "nim=?", nim).Error; err != nil {
		return err
	}

	if err := db.DbSmstr.Where("nim=?", nim).Delete(&entity.Semesteran{}).Error; err != nil {
		return err
	}

	return nil
}
