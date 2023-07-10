package handlers

import (
	"keuangan/middleware"
	"keuangan/wisuda"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandlersWsd(usecs wisuda.UsecaseWisuda, r *gin.RouterGroup) {
	eng := &useWsd{
		use: usecs,
	}

	v2 := r.Group("wisuda")
	v2.POST("", middleware.Auth(), eng.CreateWisuda)
	v2.GET("", middleware.Auth(), eng.GetAllWisuda)
	v2.GET("/:nim", middleware.AuthGetByNim(), eng.GetWisudaByNim)
	v2.PUT("/:nim", middleware.Auth(), eng.UpdateWisuda)
	v2.DELETE("/:nim", middleware.Auth(), eng.DeleteWisuda)
}

type useWsd struct {
	use wisuda.UsecaseWisuda
}

func (wsd useWsd) CreateWisuda(ctx *gin.Context) {
	err := wsd.use.CreateWisuda(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Sukses Post Wisuda"})
}

func (wsd useWsd) GetAllWisuda(ctx *gin.Context) {
	result, err := wsd.use.GetAllWisuda(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Wisuda": result, "Response": "Succes Get All Wisuda"})
}

func (wsd useWsd) GetWisudaByNim(ctx *gin.Context) {
	result, err := wsd.use.GetWisudaByNim(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (wsd useWsd) UpdateWisuda(ctx *gin.Context) {
	err := wsd.use.UpdateWisuda(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Succes Update Wisuda"})
}

func (wsd useWsd) DeleteWisuda(ctx *gin.Context) {
	err := wsd.use.DeleteWisuda(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Response": "Succes Delete Wisuda"})
}
