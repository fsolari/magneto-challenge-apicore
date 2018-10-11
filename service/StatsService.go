package service

import (
	"github.com/mercadolibre/magneto-challenge-apicore/dao"
	"github.com/mercadolibre/magneto-challenge-apicore/domain"
	"log"
	"math"
)

func CalculateDNAStats() (domain.DNAStats, error) {
	var stats domain.DNAStats
	stats, err := dao.GetDNAStats()
	if err != nil {
		log.Printf("[StatsService.CalculateDNAStats] Error getting DNA stats from DB : %s \n", err)
		return stats, err
	}
	stats.Ratio = calculateRatio(stats)

	return stats, nil
}

func calculateRatio(stats domain.DNAStats) float64 {
	if stats.CountMutantDna == 0 || stats.CountHumanDna == 0 {
		return 0
	}
	return math.Floor((float64(stats.CountHumanDna) / float64(stats.CountMutantDna)) * 100) / 100
}

/*

func setZeroIfNaN(){
	if reflect.TypeOf(stats.Ratio) == reflect.TypeOf(math.NaN()) {
		stats.Ratio = 0
	}
}
*/