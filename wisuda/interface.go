package wisuda

import (
	"keuangan/entity"

	"github.com/gin-gonic/gin"
)

type UsecaseWisuda interface {
	CreateWisuda(*gin.Context) error
	GetAllWisuda(ctx *gin.Context) ([]entity.Wisuda, error)
	GetWisudaByNim(ctx *gin.Context) ([]entity.Wisuda, error)
	UpdateWisuda(ctx *gin.Context) error
	DeleteWisuda(ctx *gin.Context) error
}

type RepositoryWisuda interface {
	CreateWisuda(entity.Wisuda, *gin.Context) error
	GetAllWisuda() ([]entity.Wisuda, error)
	GetWisudaByNim(string) ([]entity.Wisuda, error)
	UpdateWisuda(entity.Wisuda, string, *gin.Context) error
	DeleteWisuda(string) error
}
