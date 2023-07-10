package sidang

import (
	"keuangan/entity"

	"github.com/gin-gonic/gin"
)

type UsecaseSidang interface {
	CreateSidang(*gin.Context) error
	GetAllSidang(ctx *gin.Context) ([]entity.Sidang, error)
	GetSidangByNim(ctx *gin.Context) ([]entity.Sidang, error)
	UpdateSidang(ctx *gin.Context) error
	DeleteSidang(ctx *gin.Context) error
}

type RepositorySidang interface {
	CreateSidang(entity.Sidang, *gin.Context) error
	GetAllSidang() ([]entity.Sidang, error)
	GetSidangByNim(string) ([]entity.Sidang, error)
	UpdateSidang(entity.Sidang, string, *gin.Context) error
	DeleteSidang(string) error
}
