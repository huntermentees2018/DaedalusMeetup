package scheduler

import (
	"fmt"
	"log"

	"github.com/huntermentees2018/DaedalusMeetup/src/database/models"
	"github.com/jinzhu/gorm"

	"google.golang.org/api/sheets/v4"
)

// GetStudentArr returns all students in database as an array
func GetStudentArr(db *gorm.DB) []models.Student {
	rows, err := db.Model(&models.Student{}).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	students := []models.Student{}
	var student models.Student
	for rows.Next() {
		db.ScanRows(rows, &student)
		students = append(students, student)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return students
}

// Sheets does google sheets stuff
func Sheets() [][]interface{} {
	client := setupToken()

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// https://docs.google.com/spreadsheets/d/1Mz0muHRtWeSLshVHszhhizXiTIsT8vQbhWCsF8bLhnI/edit
	spreadsheetID := "1Mz0muHRtWeSLshVHszhhizXiTIsT8vQbhWCsF8bLhnI"
	readRange := "Responses!A2:J"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		return resp.Values
	}
	return nil
}
