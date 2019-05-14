package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/huntermentees2018/DaedalusMeetup/src"
	"github.com/huntermentees2018/DaedalusMeetup/src/database/models"
	"github.com/huntermentees2018/DaedalusMeetup/src/scheduler"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// PopulateDB populates the database with student data from sheets
func PopulateDB(db *gorm.DB) {
	db.DropTableIfExists(&models.Student{}, &models.Meeting{})
	db.AutoMigrate(&models.Student{}, &models.Meeting{})

	sheet := scheduler.Sheets()
	for _, row := range sheet {
		c := fmt.Sprint
		consentBool := false
		timestamp, email, name, consent, location := c(row[0]), c(row[1]), c(row[2]), c(row[3]), c(row[9])
		prefs := []models.Preference{
			models.Preference{
				Day:  "Monday",
				Time: src.DaysL(c(row[4])),
			},
			models.Preference{
				Day:  "Tuesday",
				Time: src.DaysL(c(row[5])),
			},
			models.Preference{
				Day:  "Wednesday",
				Time: src.DaysL(c(row[6])),
			},
			models.Preference{
				Day:  "Thursday",
				Time: src.DaysL(c(row[7])),
			},
			models.Preference{
				Day:  "Friday",
				Time: src.DaysL(c(row[8])),
			},
		}
		fmt.Printf("prefs = %+v\n ", prefs)
		fmt.Printf("timestamp = %+v\n ", timestamp)
		fmt.Printf("email = %+v\n ", email)
		fmt.Printf("name = %+v\n ", name)
		fmt.Printf("consent = %+v\n ", consent)
		fmt.Printf("location = %+v\n ", location)
		if consent == "Yes" {
			consentBool = true
		} else {
			consentBool = false
		}
		out, err := json.Marshal(prefs)
		if err != nil {
			panic(err)
		}
		db.Create(&models.Student{Name: name, Email: email, Consent: consentBool, Preferences: string(out), Location: location})
	}
}

//connectPostgres connects to postgres
func connectPostgres(url string, extlog bool) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	db.LogMode(true)

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

	return db
}
