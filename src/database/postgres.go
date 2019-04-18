package database

import (
	"encoding/json"
	"fmt"
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
	if extlog {
		// db.SetLogger()
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return db, nil
}

// Init sets up our database with our models
func Init() {
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Unable to read config: %v", err)
	}
	var settings map[string]string
	json.Unmarshal(b, &settings)
	db, err := connectPostgres(settings["postgresURI"], false)
	defer db.Close()
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("db: ", db)

	db.AutoMigrate(&models.Student{})
	db.Create(&models.Student{Name: "John", Email: "hpoon16@huntersoe.org"})
	var student models.Student
	db.First(&student, "Name = ?", "Miguel")
	db.Model(&student).Update("Email", "updated email")
	// db.Delete(&student)
	fmt.Println(db)
}
