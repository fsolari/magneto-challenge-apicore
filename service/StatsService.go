package service

import (
	"github.com/mercadolibre/magneto-challenge-apicore/dao"
	"github.com/mercadolibre/magneto-challenge-apicore/domain"
	"log"
	"math"
	"reflect"
)

func CalculateDNAStats() (domain.DNAStats, error) {
	var stats domain.DNAStats
	stats, err := dao.GetDNAStats()
	if err != nil {
		log.Printf("[StatsService.CalculateDNAStats] Error getting DNA stats from DB : %s \n", err)
		return stats, err
	}
	stats.Ratio = (math.Floor(CalculateRatio(stats) * 100) / 100)
	if reflect.TypeOf(stats.Ratio) == reflect.TypeOf(math.NaN()) {
		stats.Ratio = 0
	}

	return stats, nil
}

func CalculateRatio(stats domain.DNAStats) float64 {
	return float64(stats.CountHumanDna) / float64(stats.CountMutantDna)
}
