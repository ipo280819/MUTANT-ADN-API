package repositories

import "github.com/ipo280819/MUTANT-ADN-API/dto"

type MutantStatsRepository interface {
	AddStats(statsDto dto.MutantStatsDTO) error
	GetStats() (dto.MutantStatsDTO, error)
}

func NewMutantStatsRepository() MutantStatsRepository {
	return newMutantStatsRedisRepository()
}
