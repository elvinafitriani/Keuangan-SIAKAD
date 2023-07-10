package semesteran

import (
	"keuangan/entity"

	"github.com/gin-gonic/gin"
)

type UsecaseSemesteran interface {
	CreateSemesteran(*gin.Context) error
	GetAllSemesteran(ctx *gin.Context) ([]entity.Semesteran, error)
	GetSemesteranByNim(ctx *gin.Context) ([]entity.Semesteran, error)
	UpdateSemesteran(ctx *gin.Context) error
	DeleteSemesteran(ctx *gin.Context) error
}

type RepositorySemesteran interface {
	CreateSemesteran(entity.Semesteran, *gin.Context) error
	GetAllSemesteran() ([]entity.Semesteran, error)
	GetSemesteranByNim(string) ([]entity.Semesteran, error)
	UpdateSemesteran(entity.Semesteran, string, *gin.Context) error
	DeleteSemesteran(string) error
}
