package scheduler

import (
	"fmt"
	"log"
	"time"

	"github.com/huntermentees2018/DaedalusMeetup/src"
	calendar "google.golang.org/api/calendar/v3"
)

// func GetCalendarEvent(cls *calendar.EventsService) *calendar.Event {
// }

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
