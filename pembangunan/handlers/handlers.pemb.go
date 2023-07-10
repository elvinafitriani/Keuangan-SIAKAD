package handlers

import (
	"keuangan/middleware"
	"keuangan/pembangunan"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandlersPemb(usecs pembangunan.UsecasePembangunan, r *gin.RouterGroup) {
	eng := &usePemb{
		use: usecs,
	}

	v2 := r.Group("pembangunan")
	v2.POST("", middleware.Auth(), eng.CreatePembangunan)
	v2.GET("", middleware.Auth(), eng.GetAllPembangunan)
	v2.GET("/:nim", middleware.AuthGetByNim(), eng.GetPembangunanByNim)
	v2.PUT("/:nim", middleware.Auth(), eng.UpdatePembangunan)
	v2.DELETE("/:nim", middleware.Auth(), eng.DeletePembangunan)
}

type usePemb struct {
	use pembangunan.UsecasePembangunan
}

func (pemb usePemb) CreatePembangunan(ctx *gin.Context) {
	err := pemb.use.CreatePembangunan(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Sukses Post Pembangunan"})
}

func (pemb usePemb) GetAllPembangunan(ctx *gin.Context) {
	result, err := pemb.use.GetAllPembangunan(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Pembangunan": result, "Response": "Succes Get All Pembangunan"})
}

func (pemb usePemb) GetPembangunanByNim(ctx *gin.Context) {
	result, err := pemb.use.GetPembangunanByNim(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (pemb usePemb) UpdatePembangunan(ctx *gin.Context) {
	err := pemb.use.UpdatePembangunan(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Succes Update Pembangunan"})
}

func (pemb usePemb) DeletePembangunan(ctx *gin.Context) {
	err := pemb.use.DeletePembangunan(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Response": "Succes Delete Pembangunan"})
}
