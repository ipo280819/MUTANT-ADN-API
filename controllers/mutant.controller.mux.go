package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ipo280819/MUTANT-ADN-API/dto"
	"github.com/ipo280819/MUTANT-ADN-API/services"
)

type mutantMuxController struct {
	mutantService services.MutantService
}

func newMutantMuxController(mutantService services.MutantService) MutantController {
	return &mutantMuxController{mutantService}
}
func (controller *mutantMuxController) IsMutant(w http.ResponseWriter, r *http.Request) {

	var mutantDTO dto.MutantDTO
	err := json.NewDecoder(r.Body).Decode(&mutantDTO)
	if err != nil {
		ResponseErrorStatus(http.StatusBadRequest, w, err)
		return
	}

	mutant := controller.mutantService.NewMutant(mutantDTO.Dna)
	isMutant, err := mutant.IsMutant()

	if err != nil {
		ResponseError(w, err)
		return
	}

	if isMutant {
		ResponseOK(w, struct{}{})
	} else {
		ResponseStatus(http.StatusForbidden, w, struct{}{})
	}
}
