package scheduler

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/huntermentees2018/DaedalusMeetup/src"
	"github.com/huntermentees2018/DaedalusMeetup/src/database/models"
	"github.com/jinzhu/gorm"
	calendar "google.golang.org/api/calendar/v3"
)

// SchedulePeeps takes in students array and schedules
// func SchedulePeeps([]models.Student) bool {
// }

// MatchPeeps matches two students together and returns the event Id
func MatchPeeps(db *gorm.DB, cls *calendar.EventsService, stuOne models.Student, stuTwo models.Student) string {
	var location string
	rand.Seed(time.Now().UnixNano())
	choice := rand.Intn(2)
	if choice == 0 {
		location = stuOne.Location
	} else {
		location = stuTwo.Location
	}
	// TODO : create algorithm that will choose a time for the two students
	start := "2019-05-28T15:00:00-07:00"
	end := "2019-05-28T17:00:00-07:00"

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
	return resp.Id
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
