package dao

import (
	"github.com/mercadolibre/test/magneto-challenge-apicore/domain"
	"log"
)

func GetStats() (domain.Stats, error) {
	db, err := Connect()
	defer db.Close()
	var stats domain.Stats

	if err != nil {
		return stats, err
	}
	rows, err := db.Query("SELECT count(*) FROM Mutant GROUP BY MUTANT")
	if err != nil {
		return stats, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&stats.CountHumanDna, &stats.CountMutantDna)
		if err != nil {
			return stats, err
		}
		log.Println(stats.CountHumanDna, stats.CountMutantDna)
	}
	err = rows.Err()
	if err != nil {
		return stats, err
	}
	return stats, nil
}

func InsertDNA(dna domain.DNA) error {
	db, err := Connect()
	defer db.Close()

	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO Mutant(DNA, Mutant) VALUES(?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(dna.DNA, dna.Mutant)
	if err != nil {
		return err
	}
	return nil

}