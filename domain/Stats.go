package domain

type DNAStats struct {
	CountMutantDna int `json:"count_mutant_dna"`
	CountHumanDna  int `json:"count_human_dna"`
	Ratio          float64 `json:"ratio"`
}

type Count struct {
	count int
}