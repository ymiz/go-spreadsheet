package main

import (
	"fmt"
	"github.com/ymiz/go-spreadsheet/service"
	"golang.org/x/oauth2"
	"log"
)

func main() {
	config := oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://oauth2.googleapis.com/token",
		},
		RedirectURL: "urn:ietf:wg:oauth:2.0:oob",
		Scopes:      []string{"https://www.googleapis.com/auth/spreadsheets"},
	}

	srv, err := service.SheetServiceCreator{Config: &config}.Create()
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Prints the names and majors of students in a sample spreadsheet:
	spreadsheetId := ""
	readRange := "sheet1!A:A"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Name, Major:")
		for _, row := range resp.Values {
			// Print columns A and E, which correspond to indices 0 and 4.
			fmt.Printf("%s\n", row[0])
		}
	}
}
