package dao

import (
	"github.com/mercadolibre/test/magneto-challenge-apicore/domain"
	"log"
	"github.com/mercadolibre/go-meli-toolkit/goutils/logger"
)

func GetStats() (domain.Stats, error) {
	db, _ := Connect()
	defer db.Close()

	var stats domain.Stats

	tx, err := db.Begin()
	if err != nil {
		logger.Error("Failed to begin tx", err)
		return stats, err
	}

	q := "select count(ID) from mutant where Mutant = "
	err = tx.QueryRow(q + "0").Scan(&stats.CountHumanDna)
	err = tx.QueryRow(q + "1").Scan(&stats.CountMutantDna)

	if err != nil {
		logger.Error("Failed to run query", err)
		return stats, err
	}

	log.Printf("STATS INT %d %d", stats.CountHumanDna, stats.CountMutantDna)

	return stats, nil
}

func InsertDNA(dna domain.DNA) error {
	db, err := Connect()
	defer db.Close()

	if err != nil {
		log.Printf("ERROR", err)
		return err
	}

	stmt, err := db.Prepare("INSERT INTO Mutant(DNA, Mutant) VALUES(?, ?)")
	if err != nil {
		log.Printf("ERROR", err)
		return err
	}

	_, err = stmt.Exec(dna.DNA, dna.Mutant)
	if err != nil {
		log.Printf("ERROR", err)
		return err
	}
	return nil

}