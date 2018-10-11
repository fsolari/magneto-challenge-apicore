package dao

import (
	"database/sql"
	"github.com/mercadolibre/magneto-challenge-apicore/domain"
	"github.com/mercadolibre/magneto-challenge-apicore/util"
	"log"
	"reflect"
	"math"
)

func GetDNAStats() (domain.DNAStats, error) {
	db, err := Connect()
	defer db.Close()

	var stats domain.DNAStats
	var count [2]int

	if err != nil {
		log.Printf("[DNADao.GetDNAStats] Error opening DB connection: %s ", err)
		return stats, err
	}

	tx, err := db.Prepare("select IFNULL(count(ID),0) from Mutant where Mutant = ? ")
	if err != nil {
		log.Printf("[DNADao.GetDNAStats] Error preparing transaction %s", err)
		return stats, err
	}

	for i := 0; i <= 1; i++ {
		for j := range count {
			err = tx.QueryRow(string(i)).Scan(&count[j])
			switch err {
			case sql.ErrNoRows:
				count[j] = 0
			case nil:
				continue
			default:
				log.Printf("[DNADao.GetDNAStats] Error running query %s", err)
				return stats, err
			}
		}
	}

	stats.CountHumanDna = count[0]
	stats.CountMutantDna = count[1]

	return stats, nil
}

func GetDNA(dna domain.DNA) error {
	db, err := Connect()
	defer db.Close()

	var dnaString string
	if err != nil {
		log.Printf("[DNADao.GetDNARegistry] Error opening DB connection: %s ", err)
		return err
	}

	tx, err := db.Prepare("select DNA from Mutant where DNA = ? ")
	if err != nil {
		log.Printf("[DNADao.GetDNARegistry] Error preparing transaction %s", err)
		return err
	}
	row := tx.QueryRow(util.JoinStringArray(dna.DNA))

	err = row.Scan(&dnaString)
	switch err {
	case sql.ErrNoRows:
		return nil
	case nil:
		return nil
	default:
		log.Printf("[DNADao.GetDNARegistry] Error running query %s", err)
		return err
	}

	return nil
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
		log.Printf(" [DNADao.InsertDNA] Error preparing statement  : %s ", err)
		return err
	}

	dnaString := util.JoinStringArray(dna.DNA)
	_, err = stmt.Exec(dnaString, dna.Mutant)
	if err != nil {
		log.Printf("[DNADao.InsertDNA] Error executing statement : %s ", err)
		return err
	}
	return nil
}

func DeleteDNA(dna domain.DNA) error {
	db, err := Connect()
	defer db.Close()

	if err != nil {
		log.Printf("[DNADao.DeleteDNA] Error opening DB connection: %s ", err)
		return err
	}

	stmt, err := db.Prepare("DELETE FROM Mutant WHERE DNA = ?")
	if err != nil {
		log.Printf(" [DNADao.DeleteDNA] Error preparing statement  : %s ", err)
		return err
	}

	dnaString := util.JoinStringArray(dna.DNA)
	_, err = stmt.Exec(dnaString)
	if err != nil {
		log.Printf("[DNADao.DeleteDNA] Error executing statement : %s ", err)
		return err
	}
	return nil
}
