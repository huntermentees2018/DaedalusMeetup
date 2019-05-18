package main

import (
	"github.com/huntermentees2018/DaedalusMeetup/src/database"
	"github.com/huntermentees2018/DaedalusMeetup/src/scheduler"
)

// Sunday is the script that runs on sunday and matches
func Sunday() {
}

func Saturday() {
}

func main() {
	db := database.Init()
	// srv, _ := scheduler.ScheduleInit()
	_, cls := scheduler.ScheduleInit()
	defer db.Close()

	database.PopulateDB(db)
	students := scheduler.GetStudentArr(db)

	// id := scheduler.MatchPeeps(db, cls, students[0], students[1], "2019-05-28T15:00:00-07:00", "2019-05-28T17:00:00-07:00")
	// events := scheduler.GetRecent(srv)
	// for _, val := range events {
	// fmt.Println(val)
	// }

	timeTable := scheduler.CreateTimeTable(students)
	scheduler.SchedulePeeps(cls, db, timeTable)
	// m := src.LeftFromTable(timeTable)
	// fmt.Printf("m = %+v\n ", m)
	// scheduler.SchedulePeeps(db, timeTable)

	// cons := scheduler.CheckConsent(cls, id)
	// fmt.Printf("cons = %+v\n ", cons)
}
