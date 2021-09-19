package dto

type MutantStatsDTO struct {
	CountMutant int     `json:"count_mutant_dna"`
	CountHuman  int     `json:"count_human_dna"`
	Ratio       float32 `json:"ratio"`
}
