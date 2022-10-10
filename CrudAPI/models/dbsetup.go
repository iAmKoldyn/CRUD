package models

import (
	_ "fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"database/sql"
	_ "github.com/lib/pq"
)

var MPosDB *sql.DB
var MPosGORM *gorm.DB
var err error

func InitGormPostgres() {
	MPosGORM, err = gorm.Open("postgres", "user=postgres dbname=flask_db password=75yu4375 sslmode=disable")
	if err != nil {
		panic(err)
	}
}
