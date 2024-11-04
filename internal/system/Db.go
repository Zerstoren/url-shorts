package system

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type runConfig struct {
	dbHost     string
	dbPort     string
	dbUser     string
	dbPassword string
	dbName     string
	db         *gorm.DB
}

var db runConfig

func (config *runConfig) run() {
	dbCurrent, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			config.dbHost,
			config.dbUser,
			config.dbPassword,
			config.dbName,
			config.dbPort,
		), // data source name, refer https://github.com/jackc/pgx

		PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	config.db = dbCurrent
}

func GetDb() *gorm.DB {
	if db.db == nil {
		db = runConfig{
			dbHost:     "localhost",
			dbPort:     "5432",
			dbUser:     "postgres_wot",
			dbPassword: "1234",
			dbName:     "url-shorts",
		}

		db.run()
	}

	return db.db
}

type DbRequest struct {
	Db *gorm.DB
}

func (c *DbRequest) ResetDb() {
	c.Db = db.db
}

func (c *DbRequest) GetDb() *gorm.DB {
	return c.Db
}

func (c *DbRequest) SetDb(db *gorm.DB) {
	c.Db = db
}
