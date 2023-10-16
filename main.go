package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GrigorianNikc/DndCalendar/webserver"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func runWebserver() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	component := webserver.Index("Nick")
	router.GET("/test", gin.WrapH(templ.Handler(component)))
	router.Run()
}

func main() {
	file, e := os.Open("api.key")
	if e != nil {
		fmt.Println(e)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var apiKey string
	for scanner.Scan() {
		apiKey = scanner.Text()
	}

	ctx := context.Background()

	srv, err := sheets.NewService(ctx, option.WithAPIKey(apiKey))
	//srv, err := sheets.NewService(ctx, option.WithoutAuthentication())
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	spreadsheetId := ""
	readRange := "Months of Fun!B2:H7"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	mySheet, err := srv.Spreadsheets.Get(spreadsheetId).IncludeGridData(true).Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mySheet.Sheets[0].Data)
	for _, r := range mySheet.Sheets[0].Data[0].RowData {
		for _, v := range r.Values {
			fmt.Print(v.FormattedValue + ":")
			fmt.Print(v.Note)
			fmt.Print("|")
		}
		fmt.Print("\n")
	}
	t, _ := json.MarshalIndent(mySheet, "", "\t")
	//fmt.Println(string(t))
	os.WriteFile("./sheet.json", t, 0666)

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		for _, row := range resp.Values {
			for _, val := range row {
				fmt.Print(val)
				fmt.Print("|")
			}
			fmt.Print("\n")
			// Print columns A and E, which correspond to indices 0 and 4.
			//fmt.Printf("%s, %s\n", row[0], row[4])
		}
	}
	runWebserver()
}
