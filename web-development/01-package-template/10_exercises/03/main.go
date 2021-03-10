package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"
)

type Record struct {
	Date time.Time
	Open float64
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	// parse csv
	records := prs("table.csv")

	// parse template
	tpl, err := template.ParseFiles("index.tmpl")
	if err != nil {
		log.Fatalln(err)
	}

	// execute template
	err = tpl.Execute(res, records)
	if err != nil {
		log.Fatalln(err)
	}

}

func prs(filename string) []Record {
	bs, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer bs.Close()
	rdr := csv.NewReader(bs)
	rows, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	records := make([]Record, 0, len(rows))

	for i, row := range rows {
		if i == 0 {
			continue
		}
		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)

		records = append(records, Record{
			Date: date,
			Open: open,
		})
	}
	return records
}
