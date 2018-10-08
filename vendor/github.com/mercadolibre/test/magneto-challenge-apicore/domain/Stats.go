package domain

type Stats struct {
	CountMutantDna int `json:"count_mutant_dna"`
	CountHumanDna  int `json:"count_human_dna"`
	Ratio          float64 `json:"ratio"`
}