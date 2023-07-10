package handlers

import (
	"keuangan/middleware"
	"keuangan/sidang"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandlersSdg(usecs sidang.UsecaseSidang, r *gin.RouterGroup) {
	eng := &useSdg{
		use: usecs,
	}

	v2 := r.Group("sidang")
	v2.POST("", middleware.Auth(), eng.CreateSidang)
	v2.GET("", middleware.Auth(), eng.GetAllSidang)
	v2.GET("/:nim", middleware.AuthGetByNim(), eng.GetSidangByNim)
	v2.PUT("/:nim", middleware.Auth(), eng.UpdateSidang)
	v2.DELETE("/:nim", middleware.Auth(), eng.DeleteSidang)
}

type useSdg struct {
	use sidang.UsecaseSidang
}

func (sdg useSdg) CreateSidang(ctx *gin.Context) {
	err := sdg.use.CreateSidang(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Sukses Post Sidang"})
}

func (sdg useSdg) GetAllSidang(ctx *gin.Context) {
	result, err := sdg.use.GetAllSidang(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Sidang": result, "Response": "Succes Get All Sidang"})
}

func (sdg useSdg) GetSidangByNim(ctx *gin.Context) {
	result, err := sdg.use.GetSidangByNim(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (sdg useSdg) UpdateSidang(ctx *gin.Context) {
	err := sdg.use.UpdateSidang(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Succes Update Sidang"})
}

func (sdg useSdg) DeleteSidang(ctx *gin.Context) {
	err := sdg.use.DeleteSidang(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Response": "Succes Delete Sidang"})
}
