package services

import (
	"errors"
	"testing"

	"github.com/ipo280819/MUTANT-ADN-API/repositories"
	"github.com/stretchr/testify/assert"
)

func TestDefaultGetStats(t *testing.T) {
	repo := repositories.NewMutantStatsMockRepository()
	service := NewMutantStatsService(repo)
	repo.Err = errors.New("")
	statsDTO := service.GetStats()
	assert.Equal(t, 0, statsDTO.CountHuman)
	assert.Equal(t, float32(0), statsDTO.Ratio)
	assert.Equal(t, 0, statsDTO.CountMutant)
}
func TestAddStatsMutant(t *testing.T) {
	repo := repositories.NewMutantStatsMockRepository()
	service := NewMutantStatsService(repo)
	service.AddStats(true)

	statsDTO := service.GetStats()
	assert.Equal(t, 0, statsDTO.CountHuman)
	assert.Equal(t, float32(1), statsDTO.Ratio)
	assert.Equal(t, 1, statsDTO.CountMutant)
}
func TestAddStatsHuman(t *testing.T) {
	repo := repositories.NewMutantStatsMockRepository()
	service := NewMutantStatsService(repo)
	service.AddStats(false)

	statsDTO := service.GetStats()
	assert.Equal(t, 1, statsDTO.CountHuman)
	assert.Equal(t, float32(0), statsDTO.Ratio)
	assert.Equal(t, 0, statsDTO.CountMutant)
}
func TestAddStats2Human1Mutant(t *testing.T) {
	repo := repositories.NewMutantStatsMockRepository()
	service := NewMutantStatsService(repo)
	service.AddStats(false)
	service.AddStats(false)
	service.AddStats(true)

	statsDTO := service.GetStats()
	assert.Equal(t, 2, statsDTO.CountHuman)
	assert.Equal(t, float32(0.5), statsDTO.Ratio)
	assert.Equal(t, 1, statsDTO.CountMutant)
}
