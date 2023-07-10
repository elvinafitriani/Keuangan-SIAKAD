package usecase

import (
	"encoding/json"
	"errors"
	"keuangan/entity"
	"keuangan/lapbulanan"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func NewUsecaseBln(repo lapbulanan.RepositoryLapBulanan, redis *redis.Client) repoBln {
	return repoBln{
		repos:  repo,
		rediss: redis,
	}
}

type repoBln struct {
	repos  lapbulanan.RepositoryLapBulanan
	rediss *redis.Client
}

func (bln repoBln) CreateLapBulanan(ctx *gin.Context) error {
	var lapbulanan entity.LapBulanan

	if err := ctx.ShouldBindJSON(&lapbulanan); err != nil {
		return err
	}

	bulan, err := strconv.Atoi(lapbulanan.Bulan)
	if err != nil {
		return err
	}
	if bulan > 12 || bulan < 1 {
		return errors.New("periode bulan invalid")
	}

	tahun, err := strconv.Atoi(lapbulanan.Tahun)
	if err != nil {
		return err
	}
	if tahun < 2022 {
		return errors.New("periode tahun invalid")
	}

	if err := bln.repos.CreateLapBulanan(lapbulanan); err != nil {
		return err
	}

	bln.rediss.Del(ctx, "laporanbulanan")
	return nil
}

func (lapb repoBln) GetAllLapBulanan(ctx *gin.Context) ([]entity.LapBulanan, error) {
	var result []entity.LapBulanan
	var err error
	dataRedis, errRedis := lapb.rediss.Get(ctx, "laporanbulanan").Result()

	if errRedis != nil {
		result, err = lapb.repos.GetAllLapBulanan()

		if err != nil {
			return nil, err
		}

		dataJson, errJson := json.Marshal(result)

		if errJson != nil {
			return nil, errJson
		}

		err = lapb.rediss.Set(ctx, "laporanbulanan", dataJson, 0).Err()

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

func (lapb repoBln) GetLapBulananByBulan(ctx *gin.Context) ([]entity.LapBulanan, error) {
	var Inp struct {
		Bln string `uri:"bulan"`
		Thn string `uri:"tahun"`
	}

	if err := ctx.ShouldBindUri(&Inp); err != nil {
		return nil, err
	}

	result, err := lapb.repos.GetLapBulananByBulan(Inp.Bln, Inp.Thn)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (bln repoBln) DeleteLapBulanan(ctx *gin.Context) error {
	var Inp struct {
		Bln string `uri:"bulan"`
		Thn string `uri:"tahun"`
	}

	if err := ctx.ShouldBindUri(&Inp); err != nil {
		return err
	}

	err := bln.repos.DeleteLapBulanan(Inp.Bln, Inp.Thn)

	if err != nil {
		return err
	}

	bln.rediss.Del(ctx, "laporanbulanan")
	return nil
}
