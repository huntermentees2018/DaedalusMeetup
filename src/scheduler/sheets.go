package scheduler

import (
	"fmt"
	"log"

	"google.golang.org/api/sheets/v4"
)

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
		// fmt.Println("Email | Name | Yes/No | Mon | Tue | Wed | Thu | Fri | Preferred Location")
		// for _, row := range resp.Values {
		// Print columns A and E, which correspond to indices 0 and 4.
		// fmt.Printf("%s, %s, %s, %s, %s, %s, %s, %s\n", row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8])
		// }
	}
	return nil
}
