package scheduler

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/huntermentees2018/DaedalusMeetup/src"
	"github.com/huntermentees2018/DaedalusMeetup/src/database/models"
	"github.com/jinzhu/gorm"
	calendar "google.golang.org/api/calendar/v3"
)

// CreateTimeTable takes in students array and schedules
func CreateTimeTable(students []models.Student) map[string]map[rune][]uint {
	log.Println("Creating Time Table...")
	week := map[string]map[rune][]uint{
		"Monday": {
			65: []uint{},
			77: []uint{},
			69: []uint{},
		},
		"Tuesday": {
			65: []uint{},
			77: []uint{},
			69: []uint{},
		},
		"Wednesday": {
			65: []uint{},
			77: []uint{},
			69: []uint{},
		},
		"Thursday": {
			65: []uint{},
			77: []uint{},
			69: []uint{},
		},
		"Friday": {
			65: []uint{},
			77: []uint{},
			69: []uint{},
		},
	}
	for _, stu := range students {
		stuPref := []models.Preference{}
		json.Unmarshal([]byte(stu.Preferences), &stuPref)
		for _, pref := range stuPref {
			for _, c := range pref.Time {
				ids := week[pref.Day][c]
				ids = append(ids, stu.ID)
				week[pref.Day][c] = ids
			}
		}
	}
	fmt.Printf("week = %+v\n ", week)
	return week
}

// SchedulePeeps schedules the people
func SchedulePeeps(cls *calendar.EventsService, db *gorm.DB, timeTable map[string]map[rune][]uint) {
	log.Println("Scheduling people...")
	var studentOne models.Student
	var studentTwo models.Student

	isEmpty := func(ids []uint) bool { return len(ids) == 0 }
	hasPair := func(ids []uint) bool { return len(ids) > 1 }

	for src.TableIdsEach(timeTable, isEmpty) == false {
		for src.TableIdsEach(timeTable, hasPair) == false {
			for dow, day := range timeTable {
				for tod, ids := range day {
					if len(ids) > 1 {
						start, finish := src.GenerateTimeInterval(dow, tod)
						db.First(&studentOne, ids[0])
						db.First(&studentTwo, ids[1])
						log.Printf("Matching %+v and %+v", studentOne.Name, studentTwo.Name)
						MatchPeeps(db, cls, studentOne, studentTwo, start, finish)
						timeTable = src.RemoveIDFromTimeTable(timeTable, []uint{studentOne.ID, studentTwo.ID})
					}
				}
			}
		}
		left := src.LeftFromTable(timeTable)
		if len(left) == 0 {
			break
		}
		if len(left) > 1 {
			db.First(&studentOne, left[0])
			db.First(&studentTwo, left[1])
			log.Printf("Matching %+v and %+v", studentOne.Name, studentTwo.Name)
			MatchPeeps(db, cls, studentOne, studentTwo, "2019-05-28T15:00:00-07:00", "2019-05-28T17:00:00-07:00")
			timeTable = src.RemoveIDFromTimeTable(timeTable, []uint{studentOne.ID, studentTwo.ID})
		} else {
			log.Println("Somebody got left out")
			timeTable = src.RemoveIDFromTimeTable(timeTable, []uint{left[0]})
		}
	}
}

// MatchPeeps matches two students together and returns the event Id
func MatchPeeps(db *gorm.DB, cls *calendar.EventsService, stuOne models.Student, stuTwo models.Student, start string, end string) {
	var location string
	rand.Seed(time.Now().UnixNano())
	choice := rand.Intn(2)
	if choice == 0 {
		location = stuOne.Location
	} else {
		location = stuTwo.Location
	}

	event := &calendar.Event{
		Summary:         "Random Daedalus 1 on 1 Meetup!",
		Location:        location,
		Description:     src.GetAgenda(),
		GuestsCanModify: true,
		Start: &calendar.EventDateTime{
			DateTime: start,
			TimeZone: "US/Eastern",
		},
		End: &calendar.EventDateTime{
			DateTime: end,
			TimeZone: "US/Eastern",
		},
		Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=1"},
		Attendees: []*calendar.EventAttendee{
			&calendar.EventAttendee{Email: stuOne.Email},
			&calendar.EventAttendee{Email: stuTwo.Email},
		},
	}

	resp, err := cls.Insert(src.GetConfig("calendarID"), event).SendUpdates("all").Do()
	if err != nil {
		panic(err)
	}

	meeting := &models.Meeting{
		StudentOneID: stuOne.ID,
		StudentTwoID: stuTwo.ID,
		Confirmed:    false,
		EventID:      resp.Id,
		StartTime:    start,
		EndTime:      end,
	}

	db.Create(&meeting)

	log.Printf("Created meeting: %s\n", resp.Summary)
	fmt.Printf("Event created: %s\n", resp.HtmlLink)
	fmt.Printf("resp.Id = %+v\n ", resp.Id)
}

// GetRecent gets recent calendar events
func GetRecent(srv *calendar.Service) []*calendar.Event {
	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List(src.GetConfig("calendarID")).ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}

	fmt.Println("Upcoming events:")
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		return events.Items
	}

	return nil
}

// CheckConsent checks a calendar event if both parties agreed to go, if so registers to database
func CheckConsent(cls *calendar.EventsService, id string) bool {
	event, err := cls.Get(src.GetConfig("calendarID"), id).Do()
	if err != nil {
		log.Fatal(err)
	}

	return event.Attendees[0].ResponseStatus == "accepted" && event.Attendees[1].ResponseStatus == "accepted"
}

// ScheduleInit intiializes a calendar and returns both the serviec and the events service
func ScheduleInit() (*calendar.Service, *calendar.EventsService) {
	client := setupToken()

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}
	cls := calendar.NewEventsService(srv)

	return srv, cls
}
