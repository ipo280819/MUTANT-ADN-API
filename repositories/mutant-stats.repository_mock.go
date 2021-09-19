package repositories

import (
	"github.com/ipo280819/MUTANT-ADN-API/dto"
)

type MutantStatsMockRepository struct {
	Err            error
	MutantStatsDTO dto.MutantStatsDTO
}

func NewMutantStatsMockRepository() *MutantStatsMockRepository {
	return &MutantStatsMockRepository{
		Err: nil,
	}
}

func (repo *MutantStatsMockRepository) AddStats(statsDto dto.MutantStatsDTO) error {
	repo.MutantStatsDTO = statsDto
	return repo.Err
}
func (repo *MutantStatsMockRepository) GetStats() (dto.MutantStatsDTO, error) {
	return repo.MutantStatsDTO, repo.Err
}
