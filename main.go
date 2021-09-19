package main

import (
	"github.com/ipo280819/MUTANT-ADN-API/controllers"
	"github.com/ipo280819/MUTANT-ADN-API/router"
	"github.com/ipo280819/MUTANT-ADN-API/services"
)

var (
	mutantService    = services.MutantService{}
	routerHttp       = router.NewRouter()
	mutantController = controllers.NewMutantController(mutantService)
)

func main() {

	routerHttp.POST("/mutant", mutantController.IsMutant)
	routerHttp.SERVE(":3000")

}
