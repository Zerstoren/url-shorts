package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
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
			dbHost:     "127.0.0.1",
			dbPort:     "5432",
			dbUser:     os.Getenv("DB_USER"),
			dbPassword: os.Getenv("DB_PASS"),
			dbName:     os.Getenv("DB_NAME"),
		}

		db.run()
	}

	return db.db
}

type Iterable[T any] interface {
	GetFirst() (*T, bool)
}

type IterableOrigin[T any] struct {
	Origin *[]T
}

func (i *IterableOrigin[T]) GetFirst() (*T, bool) {
	if i.Origin == nil {
		return nil, false
	}

	if len(*i.Origin) == 0 {
		return nil, false
	}

	return &(*i.Origin)[0], true
}

type RequestMethods interface {
	ResetDb()
	GetDb() *gorm.DB
	SetDb(*gorm.DB)
}

type Request struct {
	Db *gorm.DB
}

func (c *Request) ResetDb() {
	c.Db = db.db
}

func (c *Request) GetDb() *gorm.DB {
	return c.Db
}

func (c *Request) SetDb(db *gorm.DB) {
	c.Db = db
}
