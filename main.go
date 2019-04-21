package main

import (
	"encoding/json"
	"fmt"

	"github.com/huntermentees2018/DaedalusMeetup/src"
	"github.com/huntermentees2018/DaedalusMeetup/src/database"
	"github.com/huntermentees2018/DaedalusMeetup/src/database/models"
	"github.com/huntermentees2018/DaedalusMeetup/src/scheduler"
	"github.com/jinzhu/gorm"
)

//PopulateDB reads form response data on sheets and populates the db
func PopulateDB(db *gorm.DB) {
	sheet := scheduler.Sheets()
	fmt.Printf("sheet = %+v\n ", sheet)
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
		// db.Model(&student).Update("Email", "updated email")
		db.Create(&models.Student{Name: name, Email: email, Consent: consentBool, Preferences: string(out), Location: location})
	}
}

func main() {
	db := database.Init()
	PopulateDB(db)
	defer db.Close()
}
