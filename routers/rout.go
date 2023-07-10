package routers

import (
	"keuangan/middleware"

	handDftr "keuangan/pendaftaran/handlers"
	repoDftr "keuangan/pendaftaran/repository"
	useDftr "keuangan/pendaftaran/usecase"

	handPemb "keuangan/pembangunan/handlers"
	repoPemb "keuangan/pembangunan/repository"
	usePemb "keuangan/pembangunan/usecase"

	handSmstr "keuangan/semesteran/handlers"
	repoSmstr "keuangan/semesteran/repository"
	useSmstr "keuangan/semesteran/usecase"

	handSdg "keuangan/sidang/handlers"
	repoSdg "keuangan/sidang/repository"
	useSdg "keuangan/sidang/usecase"

	handWsd "keuangan/wisuda/handlers"
	repoWsd "keuangan/wisuda/repository"
	useWsd "keuangan/wisuda/usecase"

	handBln "keuangan/lapbulanan/handlers"
	repoBln "keuangan/lapbulanan/repository"
	useBln "keuangan/lapbulanan/usecase"

	handThn "keuangan/laptahunan/handlers"
	repoThn "keuangan/laptahunan/repository"
	useThn "keuangan/laptahunan/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Routes struct {
	Db    *gorm.DB
	R     *gin.Engine
	Redis *redis.Client
}

func (r Routes) Routers() {
	middleware.Add(r.R, middleware.CorsMiddleware())
	v1 := r.R.Group("keuangan")

	//Pendaftaran
	repositoryDftr := repoDftr.NewRepoDftr(r.Db)
	usecaseDftr := useDftr.NewUsecaseDftr(repositoryDftr, r.Redis)
	handDftr.NewHandlersDftr(usecaseDftr, v1)

	//Pembangunan
	repositoryPemb := repoPemb.NewRepoPemb(r.Db)
	usecasePemb := usePemb.NewUsecasePemb(repositoryPemb, r.Redis)
	handPemb.NewHandlersPemb(usecasePemb, v1)

	//Semesteran
	repositorySmstr := repoSmstr.NewRepoSmstr(r.Db)
	usecaseSmstr := useSmstr.NewUsecaseSmstr(repositorySmstr, r.Redis)
	handSmstr.NewHandlersSmstr(usecaseSmstr, v1)

	//Sidang
	repositorySdg := repoSdg.NewRepoSdg(r.Db)
	usecaseSdg := useSdg.NewUsecaseSdg(repositorySdg, r.Redis)
	handSdg.NewHandlersSdg(usecaseSdg, v1)

	//Wisuda
	repositoryWsd := repoWsd.NewRepoWsd(r.Db)
	usecaseWsd := useWsd.NewUsecaseWsd(repositoryWsd, r.Redis)
	handWsd.NewHandlersWsd(usecaseWsd, v1)

	//Laporan Bulanan
	repositoryBln := repoBln.NewRepoBln(r.Db)
	usecaseBln := useBln.NewUsecaseBln(repositoryBln, r.Redis)
	handBln.NewHandlersBln(usecaseBln, v1)

	//Laporan Tahunan
	repositoryThn := repoThn.NewRepoThn(r.Db)
	usecaseThn := useThn.NewUsecaseThn(repositoryThn, r.Redis)
	handThn.NewHandlersThn(usecaseThn, v1)

}
