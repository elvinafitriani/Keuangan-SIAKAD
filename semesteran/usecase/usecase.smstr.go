package usecase

import (
	"encoding/json"
	"errors"
	"keuangan/entity"
	"keuangan/semesteran"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func NewUsecaseSmstr(repo semesteran.RepositorySemesteran, redis *redis.Client) repoSmstr {
	return repoSmstr{
		repos:  repo,
		rediss: redis,
	}
}

type repoSmstr struct {
	repos  semesteran.RepositorySemesteran
	rediss *redis.Client
}

func (smstr repoSmstr) CreateSemesteran(ctx *gin.Context) error {
	var semester entity.Semesteran

	if err := ctx.ShouldBindJSON(&semester); err != nil {
		return err
	}

	if semester.Jumlah < 100000 {
		return errors.New("jumlah invalid")
	}

	err := smstr.repos.CreateSemesteran(semester, ctx)

	if err != nil {
		return err
	}

	smstr.rediss.Del(ctx, "semester")
	return nil
}

func (smstr repoSmstr) GetAllSemesteran(ctx *gin.Context) ([]entity.Semesteran, error) {
	var result []entity.Semesteran
	var err error
	dataRedis, errRedis := smstr.rediss.Get(ctx, "semester").Result()

	if errRedis != nil {
		result, err = smstr.repos.GetAllSemesteran()

		if err != nil {
			return nil, err
		}

		dataJson, errJson := json.Marshal(result)

		if errJson != nil {
			return nil, errJson
		}

		err = smstr.rediss.Set(ctx, "semester", dataJson, 0).Err()

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

func (smstr repoSmstr) GetSemesteranByNim(ctx *gin.Context) ([]entity.Semesteran, error) {
	var Nim struct {
		Nim string `uri:"nim"`
	}

	if err := ctx.ShouldBindUri(&Nim); err != nil {
		return nil, err
	}

	result, err := smstr.repos.GetSemesteranByNim(Nim.Nim)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (smstr repoSmstr) UpdateSemesteran(ctx *gin.Context) error {
	var semester entity.Semesteran

	var Nim struct {
		Nim string `uri:"nim"`
	}
	type Semester struct {
		Nim    string `json:"nim"`
		Jumlah int    `json:"jumlah"`
	}
	var sms Semester
	if err := ctx.ShouldBindJSON(&sms); err != nil {
		return err
	}

	if err := ctx.ShouldBindUri(&Nim); err != nil {
		return err
	}

	if sms.Jumlah != 0 {
		if sms.Jumlah < 100000 {
			return errors.New("jumlah invalid")
		}
	}

	semester.Nim = sms.Nim
	semester.Jumlah = sms.Jumlah

	if err := smstr.repos.UpdateSemesteran(semester, Nim.Nim, ctx); err != nil {
		return err
	}

	smstr.rediss.Del(ctx, "semester")
	return nil

}

func (smstr repoSmstr) DeleteSemesteran(ctx *gin.Context) error {
	var Nim struct {
		Nim string `uri:"nim"`
	}

	if err := ctx.ShouldBindUri(&Nim); err != nil {
		return err
	}

	err := smstr.repos.DeleteSemesteran(Nim.Nim)

	if err != nil {
		return err
	}

	smstr.rediss.Del(ctx, "semester")
	return nil
}
