package usecase

import (
	"encoding/json"
	"errors"
	"keuangan/entity"
	"keuangan/pembangunan"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func NewUsecasePemb(repo pembangunan.RepositoryPembangunan, redis *redis.Client) repoPemb {
	return repoPemb{
		repos:  repo,
		rediss: redis,
	}
}

type repoPemb struct {
	repos  pembangunan.RepositoryPembangunan
	rediss *redis.Client
}

func (pemb repoPemb) CreatePembangunan(ctx *gin.Context) error {
	var pembangunan entity.Pembangunan

	if err := ctx.ShouldBindJSON(&pembangunan); err != nil {
		return err
	}

	if pembangunan.Jumlah < 100000 {
		return errors.New("jumlah invalid")
	}

	err := pemb.repos.CreatePembangunan(pembangunan, ctx)

	if err != nil {
		return err
	}

	pemb.rediss.Del(ctx, "pembangunan")
	return nil
}

func (pemb repoPemb) GetAllPembangunan(ctx *gin.Context) ([]entity.Pembangunan, error) {
	var result []entity.Pembangunan
	var err error
	dataRedis, errRedis := pemb.rediss.Get(ctx, "pembangunan").Result()

	if errRedis != nil {
		result, err = pemb.repos.GetAllPembangunan()

		if err != nil {
			return nil, err
		}

		dataJson, errJson := json.Marshal(result)

		if errJson != nil {
			return nil, errJson
		}

		err = pemb.rediss.Set(ctx, "pembangunan", dataJson, 0).Err()

		if err != nil {
			return nil, err
		}

		return result, nil
	}

	err = json.Unmarshal([]byte(dataRedis), &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (pemb repoPemb) GetPembangunanByNim(ctx *gin.Context) ([]entity.Pembangunan, error) {
	var Nim struct {
		Nim string `uri:"nim"`
	}

	if err := ctx.ShouldBindUri(&Nim); err != nil {
		return nil, err
	}

	result, err := pemb.repos.GetPembangunanByNim(Nim.Nim)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (pemb repoPemb) UpdatePembangunan(ctx *gin.Context) error {
	var pembangunan entity.Pembangunan

	var Nim struct {
		Nim string `uri:"nim"`
	}
	type Pemba struct {
		Nim    string `json:"nim"`
		Jumlah int    `json:"jumlah"`
	}
	var bangun Pemba

	if err := ctx.ShouldBindJSON(&bangun); err != nil {
		return err
	}

	if err := ctx.ShouldBindUri(&Nim); err != nil {
		return err
	}

	if bangun.Jumlah != 0 {
		if bangun.Jumlah < 100000 {
			return errors.New("jumlah invalid")
		}
	}

	if err := pemb.repos.UpdatePembangunan(pembangunan, Nim.Nim, ctx); err != nil {
		return err
	}

	pemb.rediss.Del(ctx, "pembangunan")
	return nil

}

func (pemb repoPemb) DeletePembangunan(ctx *gin.Context) error {
	var Nim struct {
		Nim string `uri:"nim"`
	}

	if err := ctx.ShouldBindUri(&Nim); err != nil {
		return err
	}

	err := pemb.repos.DeletePembangunan(Nim.Nim)

	if err != nil {
		return err
	}

	pemb.rediss.Del(ctx, "pembangunan")
	return nil
}
