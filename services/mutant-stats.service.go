package services

import (
	"github.com/ipo280819/MUTANT-ADN-API/dto"
	"github.com/ipo280819/MUTANT-ADN-API/repositories"
)

type MutantStatsService struct {
	repo repositories.MutantStatsRepository
}

func NewMutantStatsService(repo repositories.MutantStatsRepository) MutantStatsService {
	return MutantStatsService{
		repo,
	}
}
func (service *MutantStatsService) AddStats(isMutant bool) error {
	statsDTO := service.GetStats()
	statsDTO = calcStats(statsDTO, isMutant)
	err := service.repo.AddStats(statsDTO)
	if err != nil {
		return err
	}

	return nil
}
func (service *MutantStatsService) GetStats() dto.MutantStatsDTO {
	statsDTO, err := service.repo.GetStats()
	if err != nil {
		return dto.MutantStatsDTO{}
	}
	return statsDTO
}

func calcStats(statsDTO dto.MutantStatsDTO, isMutant bool) dto.MutantStatsDTO {
	if isMutant {
		statsDTO.CountMutant += 1
	} else {
		statsDTO.CountHuman += 1
	}
	if statsDTO.CountHuman == 0 {
		statsDTO.Ratio = 1
	} else {
		statsDTO.Ratio = float32(statsDTO.CountMutant) / float32(statsDTO.CountHuman)
	}
	return statsDTO
}
