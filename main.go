package main

import (
	"github.com/ipo280819/MUTANT-ADN-API/controllers"
	"github.com/ipo280819/MUTANT-ADN-API/repositories"
	"github.com/ipo280819/MUTANT-ADN-API/router"
	"github.com/ipo280819/MUTANT-ADN-API/services"
)

var (
	mutantStatsRepository = repositories.NewMutantStatsRepository()
	mutantService         = services.MutantService{}
	mutantStatsService    = services.NewMutantStatsService(mutantStatsRepository)
	routerHttp            = router.NewRouter()
	mutantController      = controllers.NewMutantController(mutantService, mutantStatsService)
)

func main() {

	routerHttp.POST("/mutant", mutantController.IsMutant)
	routerHttp.GET("/stats", mutantController.MutantStats)
	routerHttp.SERVE(":80")

}
