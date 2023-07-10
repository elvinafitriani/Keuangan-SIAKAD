package pembangunan

import (
	"keuangan/entity"

	"github.com/gin-gonic/gin"
)

type UsecasePembangunan interface {
	CreatePembangunan(*gin.Context) error
	GetAllPembangunan(ctx *gin.Context) ([]entity.Pembangunan, error)
	GetPembangunanByNim(ctx *gin.Context) ([]entity.Pembangunan, error)
	UpdatePembangunan(ctx *gin.Context) error
	DeletePembangunan(ctx *gin.Context) error
}

type RepositoryPembangunan interface {
	CreatePembangunan(entity.Pembangunan, *gin.Context) error
	GetAllPembangunan() ([]entity.Pembangunan, error)
	GetPembangunanByNim(string) ([]entity.Pembangunan, error)
	UpdatePembangunan(entity.Pembangunan, string, *gin.Context) error
	DeletePembangunan(string) error
}
