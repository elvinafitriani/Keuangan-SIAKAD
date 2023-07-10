package handlers

import (
	"keuangan/middleware"
	"keuangan/semesteran"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandlersSmstr(usecs semesteran.UsecaseSemesteran, r *gin.RouterGroup) {
	eng := &useSmstr{
		use: usecs,
	}

	v2 := r.Group("semesteran")
	v2.POST("", middleware.Auth(), eng.CreateSemesteran)
	v2.GET("", middleware.Auth(), eng.GetAllSemesteran)
	v2.GET("/:nim", middleware.AuthGetByNim(), eng.GetSemesteranByNim)
	v2.PUT("/:nim", middleware.Auth(), eng.UpdateSemesteran)
	v2.DELETE("/:nim", middleware.Auth(), eng.DeleteSemesteran)

}

type useSmstr struct {
	use semesteran.UsecaseSemesteran
}

func (smstr useSmstr) CreateSemesteran(ctx *gin.Context) {
	err := smstr.use.CreateSemesteran(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Sukses Post Semesteran"})
}

func (smstr useSmstr) GetAllSemesteran(ctx *gin.Context) {
	result, err := smstr.use.GetAllSemesteran(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Semesteran": result, "Response": "Succes Get All Semesteran"})
}

func (smstr useSmstr) GetSemesteranByNim(ctx *gin.Context) {
	result, err := smstr.use.GetSemesteranByNim(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (smstr useSmstr) UpdateSemesteran(ctx *gin.Context) {
	err := smstr.use.UpdateSemesteran(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Succes Update Semesteran"})
}

func (smstr useSmstr) DeleteSemesteran(ctx *gin.Context) {
	err := smstr.use.DeleteSemesteran(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Response": "Succes Delete Semesteran"})
}
