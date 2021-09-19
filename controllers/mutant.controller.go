package controllers

import (
	"net/http"

	"github.com/ipo280819/MUTANT-ADN-API/services"
)

type MutantController interface {
	IsMutant(w http.ResponseWriter, r *http.Request)
}

func NewMutantController(mutantService services.MutantService) MutantController {
	return newMutantMuxController(mutantService)
}
