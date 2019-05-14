package main

import (
	"fmt"

	"github.com/huntermentees2018/DaedalusMeetup/src/database"
	"github.com/huntermentees2018/DaedalusMeetup/src/scheduler"
)

func main() {
	db := database.Init()
	_, cls := scheduler.ScheduleInit()
	defer db.Close()

	database.PopulateDB(db)
	students := scheduler.GetStudentArr(db)

	id := scheduler.MatchPeeps(db, cls, students[0], students[1])
	cons := scheduler.CheckConsent(cls, id)
	fmt.Printf("cons = %+v\n ", cons)
}
