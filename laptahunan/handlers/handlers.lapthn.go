package handlers

import (
	"keuangan/laptahunan"
	"keuangan/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandlersThn(usecs laptahunan.UsecaseLapTahunan, r *gin.RouterGroup) {
	eng := &useThn{
		use: usecs,
	}

	v2 := r.Group("laporan-tahunan")
	v2.POST("", middleware.Auth(), eng.CreateLapTahunan)
	v2.GET("", middleware.Auth(), eng.GetAllLapTahunan)
	v2.GET("/:tahun", middleware.Auth(), eng.GetLapTahunanByTahun)
	v2.DELETE("/:tahun", middleware.Auth(), eng.DeleteLapTahunan)

}

type useThn struct {
	use laptahunan.UsecaseLapTahunan
}

func (lapt useThn) CreateLapTahunan(ctx *gin.Context) {
	err := lapt.use.CreateLapTahunan(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Sukses Post Laporan Tahunan"})
}

func (lapt useThn) GetAllLapTahunan(ctx *gin.Context) {
	result, err := lapt.use.GetAllLapTahunan(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Laporan Tahunan": result, "Response": "Succes Get All Laporan Tahunan"})
}

func (lapt useThn) GetLapTahunanByTahun(ctx *gin.Context) {
	result, err := lapt.use.GetLapTahunanByTahun(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Laporan Tahunan": result, "Response": "Succes Get Laporan Tahunan By Tahun"})
}

func (lapt useThn) DeleteLapTahunan(ctx *gin.Context) {
	err := lapt.use.DeleteLapTahunan(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Response": "Succes Delete Laporan Tahunan"})
}
