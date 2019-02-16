package config

import (
	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"os"
)

const collection string = "publishes"

//DB method
func DB() (driver.Database, error) {
	localDBHost := os.Getenv("DB_HOST")
	if localDBHost == "" {
		localDBHost = os.Getenv("ARANGO_HOST")
	}
	init, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{localDBHost},
	})
	con, err := init.SetAuthentication(driver.BasicAuthentication(os.Getenv("ARANGO_USERNAME"), ""));
	if err != nil {
		return nil, err
	}
	c, err := driver.NewClient(driver.ClientConfig{
		Connection: con,
	})
	if err != nil {
		return nil, err
	}
	db, err := c.Database(nil, os.Getenv("ARANGO_DATABASE"))
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
