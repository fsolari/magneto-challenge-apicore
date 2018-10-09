package dao

import (
	"github.com/mercadolibre/magneto-challenge-apicore/domain"
	"log"
	"strings"
)

func GetDNAStats() (domain.DNAStats, error) {
	db, err := Connect()
	defer db.Close()

	var stats domain.DNAStats

	if err != nil {
		log.Printf("[DNADao.GetDNAStats] Error opening DB connection: %s ", err)
		return stats, err
	}

	tx, err := db.Prepare("select count(ID) from Mutant where Mutant = ? ")
	if err != nil {
		log.Printf("[DNADao.GetDNAStats] Error preparing transaction %s", err)
		return stats, err
	}

	err = tx.QueryRow("0").Scan(&stats.CountHumanDna)
	err = tx.QueryRow("1").Scan(&stats.CountMutantDna)

	if err != nil {
		log.Printf("[DNADao.GetDNAStats] Error running query %s", err)
		return stats, err
	}

	return stats, nil
}

func InsertDNA(dna domain.DNA) error {
	db, err := Connect()
	defer db.Close()

	if err != nil {
		log.Printf("[DNADao.InsertDNA] Error opening DB connection: %s ", err)
		return err
	}

	stmt, err := db.Prepare("INSERT INTO Mutant(DNA, Mutant) VALUES(?, ?)")
	if err != nil {
		log.Printf(" [DNADao.InsertDNA] Error preparing statement B : %s ", err)
		return err
	}

	dnaString := strings.Join(dna.DNA, " ")
	_, err = stmt.Exec(dnaString, dna.Mutant)
	if err != nil {
		log.Printf("[DNADao.InsertDNA] Error executing statement : %s ", err)
		return err
	}
	return nil

}