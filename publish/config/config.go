package config

import (
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"os"
)

const collection string = "publishes"

//DB method
func DB() (driver.Database, error) {
	localDBHost := os.Getenv("DB_HOST")
	if localDBHost == "" {
		localDBHost = os.Getenv("ARANGODB_HOST")
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

//Collection context
func Collection() (driver.Collection, error) {
	db, err := DB()
	if err != nil {
		return nil, err
	}
	c, err := db.Collection(nil, collection)
	if err != nil {
		return nil, err
	}
	return c, nil
}
