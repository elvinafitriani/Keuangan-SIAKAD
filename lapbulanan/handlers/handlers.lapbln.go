package handlers

import (
	"keuangan/lapbulanan"
	"keuangan/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHandlersBln(usecs lapbulanan.UsecaseLapBulanan, r *gin.RouterGroup) {
	eng := &useBln{
		use: usecs,
	}

	v2 := r.Group("laporan-bulanan")
	v2.POST("", middleware.Auth(), eng.CreateLapBulanan)
	v2.GET("", middleware.Auth(), eng.GetAllLapBulanan)
	v2.GET("/:bulan/:tahun", middleware.Auth(), eng.GetLapBulananByBulan)
	v2.DELETE("/:bulan/:tahun", middleware.Auth(), eng.DeleteLapBulanan)

}

type useBln struct {
	use lapbulanan.UsecaseLapBulanan
}

func (lapb useBln) CreateLapBulanan(ctx *gin.Context) {
	err := lapb.use.CreateLapBulanan(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Message": "Sukses Post Laporan Bulanan"})
}

func (lapb useBln) GetAllLapBulanan(ctx *gin.Context) {
	result, err := lapb.use.GetAllLapBulanan(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Laporan Bulanan": result, "Response": "Succes Get All Laporan Bulanan"})
}

func (lapb useBln) GetLapBulananByBulan(ctx *gin.Context) {
	result, err := lapb.use.GetLapBulananByBulan(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Laporan Bulanan": result, "Response": "Succes Get Laporan Bulanan By Periode Bulan"})
}

func (lapb useBln) DeleteLapBulanan(ctx *gin.Context) {
	err := lapb.use.DeleteLapBulanan(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"Response": "Succes Delete Laporan Bulanan"})
}
