package main

import (
	"encoding/json"
	"fmt"
	"log"

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

// func MatchPeeps(db *gorm.DB, models.Student, models.Student) {
// event := &calendar.Event{
// 	Summary:         "Google I/O 2015",
// 	Location:        "800 Howard St., San Francisco, CA 94103",
// 	Description:     "A chance to hear more about Google's developer products.",
// 	GuestsCanModify: true,
// 	Start: &calendar.EventDateTime{
// 		DateTime: "2019-05-28T09:00:00-07:00",
// 		TimeZone: "America/Los_Angeles",
// 	},
// 	End: &calendar.EventDateTime{
// 		DateTime: "2019-05-28T17:00:00-07:00",
// 		TimeZone: "America/Los_Angeles",
// 	},
// 	Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=1"},
// 	Attendees: []*calendar.EventAttendee{
// 		&calendar.EventAttendee{Email: "miguelacero528@gmail.com"},
// 		&calendar.EventAttendee{Email: "macero16@huntersoe.org"},
// 	},
// }
// event, _ = cls.Insert(src.GetConfig("calendarID"), event).SendUpdates("all").Do()
// fmt.Printf("Event created: %s\n", event.HtmlLink)
// fmt.Println(event.Id)
// }

func main() {
	db := database.Init()
	PopulateDB(db)
	rows, err := db.Find(&models.Student{}).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var students []models.Student
	var student models.Student
	for rows.Next() {
		rows.Scan(&student.Name, &student.Email, &student.Consent, &student.Preferences, &student.Location)
		students = append(students, student)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(students)
	// MatchPeeps
	defer db.Close()
	// _, cls := scheduler.ScheduleInit()
	// events := scheduler.GetRecent(srv)

	// event, err := cls.Get(src.GetConfig("calendarID"), "72jfrbvosa6mmcvu112qokp1pg").Do()
	// if err != nil {
	// panic(err)
	// }
	// fmt.Println(event)

}
