package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"log"
	"strings"
	"fmt"
	"net/url"
	"github.com/mercadolibre/magneto-challenge-apicore/domain"
	"net"
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
	var config domain.Configuration

	u, err := url.Parse(os.Getenv("JAWSDB_URL"))
	if err != nil {
		panic(err)
	}
	config.Engine = u.Scheme
	config.User = u.User.Username()
	config.Password, _ = u.User.Password()
	config.Server, config.Port, _ = net.SplitHostPort(u.Host)
	config.Database = strings.Trim(u.Path, "/")

	return config, nil
}
