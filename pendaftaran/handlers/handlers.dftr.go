package handlers

import (
	"keuangan/middleware"
	"keuangan/pendaftaran"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandlersDftr(usecs pendaftaran.UsecasePendaftaran, r *gin.RouterGroup) {
	eng := &useDftr{
		use: usecs,
	}

	v2 := r.Group("pendaftaran")
	v2.POST("", middleware.Auth(), eng.CreatePendaftaran)
	v2.GET("", middleware.Auth(), eng.GetAllPendaftaran)
	v2.GET("/:nim", middleware.AuthGetByNim(), eng.GetPendaftaranByNim)
	v2.PUT("/:nim", middleware.Auth(), eng.UpdatePendaftaran)
	v2.DELETE("/:nim", middleware.Auth(), eng.DeletePendaftaran)
}

type useDftr struct {
	use pendaftaran.UsecasePendaftaran
}

func (dftr useDftr) CreatePendaftaran(ctx *gin.Context) {
	err := dftr.use.CreatePendaftaran(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Sukses Post Pendaftaran"})
}

func (dftr useDftr) GetAllPendaftaran(ctx *gin.Context) {
	result, err := dftr.use.GetAllPendaftaran(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Pendaftaran": result, "Response": "Succes Get All Pendaftaran"})
}

func (dftr useDftr) GetPendaftaranByNim(ctx *gin.Context) {
	result, err := dftr.use.GetPendaftaranByNim(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (dftr useDftr) UpdatePendaftaran(ctx *gin.Context) {
	err := dftr.use.UpdatePendaftaran(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Succes Update Pendaftaran"})
}

func (dftr useDftr) DeletePendaftaran(ctx *gin.Context) {
	err := dftr.use.DeletePendaftaran(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Response": "Succes Delete Pendaftaran"})
}
