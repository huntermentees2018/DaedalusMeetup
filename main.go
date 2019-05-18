package main

import (
	"github.com/huntermentees2018/DaedalusMeetup/src/database"
	"github.com/huntermentees2018/DaedalusMeetup/src/scheduler"
)

// Sunday is the script that runs on sunday and matches people to certain days
func Sunday() {
}

// Saturday looks at all the event Ids inside the DB and checks whether people actually met up and updates db accordingly.
func Saturday() {
	// cons := scheduler.CheckConsent(cls, id)
	// fmt.Printf("cons = %+v\n ", cons)
}

func main() {
	db := database.Init()
	// srv, _ := scheduler.ScheduleInit()
	_, cls := scheduler.ScheduleInit()
	defer db.Close()

	database.PopulateDB(db)
	students := scheduler.GetStudentArr(db)

	timeTable := scheduler.CreateTimeTable(students)
	scheduler.SchedulePeeps(cls, db, timeTable)

	// events := scheduler.GetRecent(srv)
	// for _, val := range events {
	// fmt.Println(val)
	// }
}
