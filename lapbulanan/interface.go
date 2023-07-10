package lapbulanan

import (
	"keuangan/entity"

	"github.com/gin-gonic/gin"
)

type UsecaseLapBulanan interface {
	CreateLapBulanan(*gin.Context) error
	GetAllLapBulanan(ctx *gin.Context) ([]entity.LapBulanan, error)
	GetLapBulananByBulan(ctx *gin.Context) ([]entity.LapBulanan, error)
	DeleteLapBulanan(ctx *gin.Context) error
}

type RepositoryLapBulanan interface {
	CreateLapBulanan(entity.LapBulanan) error
	GetAllLapBulanan() ([]entity.LapBulanan, error)
	GetLapBulananByBulan(string, string) ([]entity.LapBulanan, error)
	DeleteLapBulanan(string, string) error
}
