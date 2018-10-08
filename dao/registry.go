package dao

import (
	"github.com/mercadolibre/magneto-challenge-apicore/domain"
	"log"
)

func GetStats() (domain.Stats, error) {
	db, _ := Connect()
	defer db.Close()

	var stats domain.Stats

	tx, err := db.Prepare("select count(ID) from Mutant where Mutant = ? ")
	if err != nil {
		log.Printf("Failed to prepare query %s", err)
		return stats, err
	}

	err = tx.QueryRow("0").Scan(&stats.CountHumanDna)
	err = tx.QueryRow("1").Scan(&stats.CountMutantDna)

	if err != nil {
		log.Printf("Failed to run query %s", err)
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