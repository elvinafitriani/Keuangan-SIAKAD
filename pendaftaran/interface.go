package pendaftaran

import (
	"keuangan/entity"

	"github.com/gin-gonic/gin"
)

type UsecasePendaftaran interface {
	CreatePendaftaran(*gin.Context) error
	GetAllPendaftaran(ctx *gin.Context) ([]entity.Pendaftaran, error)
	GetPendaftaranByNim(ctx *gin.Context) ([]entity.Pendaftaran, error)
	UpdatePendaftaran(ctx *gin.Context) error
	DeletePendaftaran(ctx *gin.Context) error
}

type RepositoryPendaftaran interface {
	CreatePendaftaran(entity.Pendaftaran, *gin.Context) error
	GetAllPendaftaran() ([]entity.Pendaftaran, error)
	GetPendaftaranByNim(string) ([]entity.Pendaftaran, error)
	UpdatePendaftaran(entity.Pendaftaran, string, *gin.Context) error
	DeletePendaftaran(string) error
}
