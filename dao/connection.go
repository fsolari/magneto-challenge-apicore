package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"os"
	"github.com/mercadolibre/test/magneto-challenge-apicore/domain"
	"encoding/json"
	"log"
)

func Connect() (*sql.DB, error) {
	config, err := getConfiguration()
	if err != nil {
		log.Printf("ERROR", err)
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true", config.User, config.Password, config.Server, config.Port, config.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("ERROR", err)
		return nil, err
	}
	return db, nil
}

func getConfiguration() (domain.Configuration, error) {
	config := domain.Configuration{}
	file, err := os.Open("./config.json")
	if err != nil {
		log.Printf("ERROR", err)
		return config, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Printf("ERROR", err)
		return config, err
	}

	return config, nil
}
