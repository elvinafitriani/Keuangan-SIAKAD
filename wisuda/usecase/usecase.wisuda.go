package usecase

import (
	"encoding/json"
	"errors"
	"keuangan/entity"
	"keuangan/wisuda"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func NewUsecaseWsd(repo wisuda.RepositoryWisuda, redis *redis.Client) repoWsd {
	return repoWsd{
		repos:  repo,
		rediss: redis,
	}
}

type repoWsd struct {
	repos  wisuda.RepositoryWisuda
	rediss *redis.Client
}

func (wsd repoWsd) CreateWisuda(ctx *gin.Context) error {
	var wisuda entity.Wisuda

	if err := ctx.ShouldBindJSON(&wisuda); err != nil {
		return err
	}

	if wisuda.Jumlah < 100000 {
		return errors.New("jumlah invalid")
	}

	err := wsd.repos.CreateWisuda(wisuda, ctx)

	if err != nil {
		return err
	}

	wsd.rediss.Del(ctx, "wisuda")
	return nil
}

func (wsd repoWsd) GetAllWisuda(ctx *gin.Context) ([]entity.Wisuda, error) {
	var result []entity.Wisuda
	var err error
	dataRedis, errRedis := wsd.rediss.Get(ctx, "wisuda").Result()

	if errRedis != nil {
		result, err = wsd.repos.GetAllWisuda()

		if err != nil {
			return nil, err
		}

		dataJson, errJson := json.Marshal(result)

		if errJson != nil {
			return nil, errJson
		}

		err = wsd.rediss.Set(ctx, "wisuda", dataJson, 0).Err()

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

func (wsd repoWsd) GetWisudaByNim(ctx *gin.Context) ([]entity.Wisuda, error) {
	var Nim struct {
		Nim string `uri:"nim"`
	}

	if err := ctx.ShouldBindUri(&Nim); err != nil {
		return nil, err
	}

	result, err := wsd.repos.GetWisudaByNim(Nim.Nim)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (wsd repoWsd) UpdateWisuda(ctx *gin.Context) error {
	var wisuda entity.Wisuda
	type Wisuda struct {
		Nim    string `json:"nim"`
		Jumlah int    `json:"jumlah"`
	}
	var wis Wisuda

	var Nim struct {
		Nim string `uri:"nim"`
	}

	if err := ctx.ShouldBindJSON(&wis); err != nil {
		return err
	}

	if err := ctx.ShouldBindUri(&Nim); err != nil {
		return err
	}

	if wis.Jumlah != 0 {
		if wis.Jumlah < 100000 {
			return errors.New("jumlah invalid")
		}
	}

	wisuda.Nim = wis.Nim
	wisuda.Jumlah = wis.Jumlah

	if err := wsd.repos.UpdateWisuda(wisuda, Nim.Nim, ctx); err != nil {
		return err
	}

	wsd.rediss.Del(ctx, "wisuda")
	return nil

}

func (wsd repoWsd) DeleteWisuda(ctx *gin.Context) error {
	var Nim struct {
		Nim string `uri:"nim"`
	}

	if err := ctx.ShouldBindUri(&Nim); err != nil {
		return err
	}

	err := wsd.repos.DeleteWisuda(Nim.Nim)

	if err != nil {
		return err
	}

	wsd.rediss.Del(ctx, "wisuda")
	return nil
}
