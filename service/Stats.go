package service

import (
	"github.com/mercadolibre/test/magneto-challenge-apicore/domain"
	"github.com/mercadolibre/test/magneto-challenge-apicore/dao"
)

func CalculateStats() (domain.Stats, error) {
	var stats domain.Stats
	stats, err := dao.GetStats()
	if err != nil {
		return stats, err
	}
	stats.Ratio = CalculateRatio(stats)
	return stats, nil
}

func CalculateRatio(stats domain.Stats) float64 {
	return float64(stats.CountHumanDna) / float64(stats.CountMutantDna)
}
