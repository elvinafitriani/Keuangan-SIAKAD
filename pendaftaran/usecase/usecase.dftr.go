package usecase

import (
	"encoding/json"
	"errors"
	"keuangan/entity"
	"keuangan/pendaftaran"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func NewUsecaseDftr(repo pendaftaran.RepositoryPendaftaran, redis *redis.Client) repoDftr {
	return repoDftr{
		repos:  repo,
		rediss: redis,
	}
}

type repoDftr struct {
	repos  pendaftaran.RepositoryPendaftaran
	rediss *redis.Client
}

func (dftr repoDftr) CreatePendaftaran(ctx *gin.Context) error {
	var daftar entity.Pendaftaran

	if err := ctx.ShouldBindJSON(&daftar); err != nil {
		return err
	}

	if daftar.Jumlah < 100000 {
		return errors.New("jumlah invalid")
	}

	err := dftr.repos.CreatePendaftaran(daftar, ctx)

	if err != nil {
		return err
	}

	dftr.rediss.Del(ctx, "daftar")
	return nil
}

func (dftr repoDftr) GetAllPendaftaran(ctx *gin.Context) ([]entity.Pendaftaran, error) {
	var result []entity.Pendaftaran
	var err error
	dataRedis, errRedis := dftr.rediss.Get(ctx, "daftar").Result()

	if errRedis != nil {
		result, err = dftr.repos.GetAllPendaftaran()

		if err != nil {
			return nil, err
		}

		dataJson, errJson := json.Marshal(result)

		if errJson != nil {
			return nil, errJson
		}

		err = dftr.rediss.Set(ctx, "daftar", dataJson, 0).Err()

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

func (dftr repoDftr) GetPendaftaranByNim(ctx *gin.Context) ([]entity.Pendaftaran, error) {
	var Nim struct {
		Nim string `uri:"nim"`
	}

	if err := ctx.ShouldBindUri(&Nim); err != nil {
		return nil, err
	}

	result, err := dftr.repos.GetPendaftaranByNim(Nim.Nim)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (dftr repoDftr) UpdatePendaftaran(ctx *gin.Context) error {
	var daftar entity.Pendaftaran
	type Pendaftaran struct {
		Nim    string `json:"nim"`
		Jumlah int    `json:"jumlah"`
	}
	var Nim struct {
		Nim string `uri:"nim"`
	}

	var pend Pendaftaran

	if err := ctx.ShouldBindJSON(&pend); err != nil {
		return err
	}

	if err := ctx.ShouldBindUri(&Nim); err != nil {
		return err
	}

	if pend.Jumlah != 0 {
		if pend.Jumlah < 100000 {
			return errors.New("jumlah invalid")
		}
	}

	daftar.Nim = pend.Nim
	daftar.Jumlah = pend.Jumlah

	if err := dftr.repos.UpdatePendaftaran(daftar, Nim.Nim, ctx); err != nil {
		return err
	}

	dftr.rediss.Del(ctx, "daftar")
	return nil

}

func (dftr repoDftr) DeletePendaftaran(ctx *gin.Context) error {
	var Nim struct {
		Nim string `uri:"nim"`
	}

	if err := ctx.ShouldBindUri(&Nim); err != nil {
		return err
	}

	err := dftr.repos.DeletePendaftaran(Nim.Nim)

	if err != nil {
		return err
	}

	dftr.rediss.Del(ctx, "daftar")
	return nil
}
