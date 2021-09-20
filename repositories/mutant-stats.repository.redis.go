package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/ipo280819/MUTANT-ADN-API/dto"
)

const MUTANT_STATS_KEY = "mutant_stats"

type mutantStatsRedisRepository struct {
	client *redis.Client
}

// NewRepository instances a Redis implementation of the gopherapi.Repository
func newMutantStatsRedisRepository() *mutantStatsRedisRepository {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis-service:6379",
		Password: os.Getenv("REDIS_PASS"),
		DB:       0, // use default DB
	})
	ctx := context.Background()
	status, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(status)
	return &mutantStatsRedisRepository{
		client,
	}
}

func (repo *mutantStatsRedisRepository) AddStats(statsDto dto.MutantStatsDTO) error {
	ctx := context.Background()

	stats, err := json.Marshal(statsDto)
	if err != nil {
		return err
	}

	errR := repo.client.Set(ctx, MUTANT_STATS_KEY, string(stats), 0)
	if err != nil {
		return errR.Err()
	}
	return nil
}
func (repo *mutantStatsRedisRepository) GetStats() (dto.MutantStatsDTO, error) {
	ctx := context.Background()

	val, err := repo.client.Get(ctx, MUTANT_STATS_KEY).Result()

	if err != nil {
		return dto.MutantStatsDTO{}, err
	}

	var statsDTO dto.MutantStatsDTO

	err = json.Unmarshal([]byte(val), &statsDTO)

	if err != nil {
		return dto.MutantStatsDTO{}, err
	}
	return statsDTO, err
}
