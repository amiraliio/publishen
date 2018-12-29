package config

import (
	driver "github.com/arangodb/go-driver"
	http "github.com/arangodb/go-driver/http"
	os "os"
)

var DBHost = os.Getenv("ARANGODB_HOST")

func DB() (driver.Database, error) {

	localDBHost := os.Getenv("DB_HOST")
	if localDBHost == "" {
		localDBHost = DBHost
	}

	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{localDBHost},
	})
	if err != nil {
		return nil, err
	}

	c, err := driver.NewClient(driver.ClientConfig{
		Connection: conn,
	})
	if err != nil {
		return nil, err
	}

	db, err := c.Database(nil, os.Getenv("DATABASE_NAME"))
	if err != nil {
		return nil, err
	}
	return db, nil
}
