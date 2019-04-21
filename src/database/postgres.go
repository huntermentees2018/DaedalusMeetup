package database

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/huntermentees2018/DaedalusMeetup/src/database/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//connectPostgres connects to postgres
func connectPostgres(url string, extlog bool) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	db.LogMode(true)
	//external logging file
	if extlog {
		// db.SetLogger(log.New(f, ))
	}

	return db, nil
}

// Init sets up our database with our models
func Init() *gorm.DB {
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Unable to read config: %v", err)
	}
	var settings map[string]string
	json.Unmarshal(b, &settings)
	db, err := connectPostgres(settings["postgresURI"], false)

	if err != nil {
		panic(err)
	}

	db.DropTableIfExists(&models.Student{}, &models.Meeting{})
	db.AutoMigrate(&models.Student{}, &models.Meeting{})

	return db
}
