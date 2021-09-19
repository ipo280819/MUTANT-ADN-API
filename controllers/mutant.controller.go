package controllers

import (
	"net/http"

	"github.com/ipo280819/MUTANT-ADN-API/services"
)

type MutantController interface {
	IsMutant(w http.ResponseWriter, r *http.Request)
	MutantStats(w http.ResponseWriter, r *http.Request)
}

func NewMutantController(mutantService services.MutantService, mutantStatsService services.MutantStatsService) MutantController {
	return newMutantMuxController(mutantService, mutantStatsService)
}
