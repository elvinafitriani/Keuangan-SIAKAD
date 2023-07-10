package usecase

import (
	"encoding/json"
	"errors"
	"keuangan/entity"
	"keuangan/sidang"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func NewUsecaseSdg(repo sidang.RepositorySidang, redis *redis.Client) repoSdg {
	return repoSdg{
		repos:  repo,
		rediss: redis,
	}
}

type repoSdg struct {
	repos  sidang.RepositorySidang
	rediss *redis.Client
}

func (sdg repoSdg) CreateSidang(ctx *gin.Context) error {
	var sidang entity.Sidang

	if err := ctx.ShouldBindJSON(&sidang); err != nil {
		return err
	}

	if sidang.Jumlah < 100000 {
		return errors.New("jumlah invalid")
	}

	err := sdg.repos.CreateSidang(sidang, ctx)

	if err != nil {
		return err
	}

	sdg.rediss.Del(ctx, "sidang")
	return nil
}

func (sdg repoSdg) GetAllSidang(ctx *gin.Context) ([]entity.Sidang, error) {
	var result []entity.Sidang
	var err error
	dataRedis, errRedis := sdg.rediss.Get(ctx, "sidang").Result()

	if errRedis != nil {
		result, err = sdg.repos.GetAllSidang()

		if err != nil {
			return nil, err
		}

		dataJson, errJson := json.Marshal(result)

		if errJson != nil {
			return nil, errJson
		}

		err = sdg.rediss.Set(ctx, "sidang", dataJson, 0).Err()

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

func (sdg repoSdg) GetSidangByNim(ctx *gin.Context) ([]entity.Sidang, error) {
	var Nim struct {
		Nim string `uri:"nim"`
	}

	if err := ctx.ShouldBindUri(&Nim); err != nil {
		return nil, err
	}

	result, err := sdg.repos.GetSidangByNim(Nim.Nim)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (sdg repoSdg) UpdateSidang(ctx *gin.Context) error {
	var sidang entity.Sidang

	var Nim struct {
		Nim string `uri:"nim"`
	}
	type Sidang struct {
		Nim    string `json:"nim"`
		Jumlah int    `json:"jumlah"`
	}
	var sid Sidang

	if err := ctx.ShouldBindJSON(&sid); err != nil {
		return err
	}

	if err := ctx.ShouldBindUri(&Nim); err != nil {
		return err
	}

	if sid.Jumlah != 0 {
		if sid.Jumlah < 100000 {
			return errors.New("jumlah invalid")
		}
	}

	sidang.Nim = sid.Nim
	sidang.Jumlah = sid.Jumlah

	if err := sdg.repos.UpdateSidang(sidang, Nim.Nim, ctx); err != nil {
		return err
	}

	sdg.rediss.Del(ctx, "sidang")
	return nil

}

func (sdg repoSdg) DeleteSidang(ctx *gin.Context) error {
	var Nim struct {
		Nim string `uri:"nim"`
	}

	if err := ctx.ShouldBindUri(&Nim); err != nil {
		return err
	}

	err := sdg.repos.DeleteSidang(Nim.Nim)

	if err != nil {
		return err
	}

	sdg.rediss.Del(ctx, "sidang")
	return nil
}
