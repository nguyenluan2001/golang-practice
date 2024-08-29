package database

import (
	"fmt"
	"log"

	"github.com/nguyenluan2001/golang-authenticate/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func (psql *Postgres) Connect() {
	dsn := "host=localhost user=postgres password=luandatabase dbname=go-authenticate port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Database connection failed")
	}
	fmt.Println("Database connection success")
	psql.DB = db
}
func (psql *Postgres) Migrate() {
	psql.DB.AutoMigrate(&model.UserSchema{})
	psql.DB.AutoMigrate(&model.TodoSchema{})
}
