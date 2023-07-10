package laptahunan

import (
	"keuangan/entity"

	"github.com/gin-gonic/gin"
)

type UsecaseLapTahunan interface {
	CreateLapTahunan(*gin.Context) error
	GetAllLapTahunan(ctx *gin.Context) ([]entity.LapTahunan, error)
	GetLapTahunanByTahun(ctx *gin.Context) ([]entity.LapTahunan, error)
	DeleteLapTahunan(ctx *gin.Context) error
}

type RepositoryLapTahunan interface {
	CreateLapTahunan(entity.LapTahunan) error
	GetAllLapTahunan() ([]entity.LapTahunan, error)
	GetLapTahunanByTahun(string) ([]entity.LapTahunan, error)
	DeleteLapTahunan(string) error
}
