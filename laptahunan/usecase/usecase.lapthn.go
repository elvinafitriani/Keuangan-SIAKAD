package usecase

import (
	"encoding/json"
	"errors"
	"keuangan/entity"
	"keuangan/laptahunan"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func NewUsecaseThn(repo laptahunan.RepositoryLapTahunan, redis *redis.Client) repoThn {
	return repoThn{
		repos:  repo,
		rediss: redis,
	}
}

type repoThn struct {
	repos  laptahunan.RepositoryLapTahunan
	rediss *redis.Client
}

func (thn repoThn) CreateLapTahunan(ctx *gin.Context) error {
	var laptahunan entity.LapTahunan

	if err := ctx.ShouldBindJSON(&laptahunan); err != nil {
		return err
	}

	tahun, err := strconv.Atoi(laptahunan.Tahun)
	if err != nil {
		return err
	}

	if tahun < 2022 {
		return errors.New("periode tahun invalid")
	}

	if err := thn.repos.CreateLapTahunan(laptahunan); err != nil {
		return err
	}

	thn.rediss.Del(ctx, "laporantahunan")
	return nil
}

func (lapt repoThn) GetAllLapTahunan(ctx *gin.Context) ([]entity.LapTahunan, error) {
	var result []entity.LapTahunan
	var err error
	dataRedis, errRedis := lapt.rediss.Get(ctx, "laporantahunan").Result()

	if errRedis != nil {
		result, err = lapt.repos.GetAllLapTahunan()

		if err != nil {
			return nil, err
		}

		dataJson, errJson := json.Marshal(result)

		if errJson != nil {
			return nil, errJson
		}

		err = lapt.rediss.Set(ctx, "laporantahunan", dataJson, 0).Err()

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

func (lapt repoThn) GetLapTahunanByTahun(ctx *gin.Context) ([]entity.LapTahunan, error) {
	var Inp struct {
		Thn string `uri:"tahun"`
	}

	if err := ctx.ShouldBindUri(&Inp); err != nil {
		return nil, err
	}

	result, err := lapt.repos.GetLapTahunanByTahun(Inp.Thn)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (thn repoThn) DeleteLapTahunan(ctx *gin.Context) error {
	var Tahun struct {
		Thn string `uri:"tahun"`
	}

	if err := ctx.ShouldBindUri(&Tahun); err != nil {
		return err
	}

	err := thn.repos.DeleteLapTahunan(Tahun.Thn)

	if err != nil {
		return err
	}

	thn.rediss.Del(ctx, "laporantahunan")
	return nil
}
