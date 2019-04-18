package config

import (
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/micro/go-config"
	"os"
)

//DB method
func DB() (driver.Database, error) {
	localDBHost := os.Getenv("DB_HOST")
	if localDBHost == "" {
		localDBHost = config.Get("database", "arango", "host").String("localhost")
	}
	init, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{localDBHost},
	})
	con, err := init.SetAuthentication(driver.BasicAuthentication(config.Get("database", "arango", "username").String("root"), config.Get("database", "arango", "password").String("")))
	if err != nil {
		return nil, err
	}
	c, err := driver.NewClient(driver.ClientConfig{
		Connection: con,
	})
	if err != nil {
		return nil, err
	}
	db, err := c.Database(nil, config.Get("database", "arango", "dbName").String("publishen"))
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
	c, err := db.Collection(nil, config.Get("database", "arango", "collection").String("publishes"))
	if err != nil {
		return nil, err
	}
	return c, nil
}
