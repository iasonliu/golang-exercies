package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type thumbnail struct {
	URL           string
	Height, Width int
}

type img struct {
	Width, Height int
	Title         string
	Thumbnail     thumbnail
	Animated      bool
	IDs           []int
}

// You can choose to only unmarshal some of the json data
// Create a data structure that only has fields for some of the data
type city struct {
	Latitude, Longitude float64
	City                string
}

type cities []city

func main() {
	var data img
	rcvd := `{"Width":800,"Height":600,"Title":"View from 15th Floor","Thumbnail":{"Url":"http://www.example.com/image/481989943","Height":125,"Width":100},"Animated":false,"IDs":[116,943,234,38793]}`

	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln("error unmarshalling", err)
	}
	fmt.Println(data)

	for i, v := range data.IDs {
		fmt.Println(i, v)
	}
	fmt.Println(data.Thumbnail.URL)
	fmt.Println("======================")
	var data1 cities
	rcvd1 := `[{"precision":"zip","Latitude":37.7668,"Longitude":-122.3959,"Address":"","City":"SAN FRANCISCO","State":"CA","Zip":"94107","Country":"US"},{"precision":"zip","Latitude":37.371991,"Longitude":-122.02602,"Address":"","City":"SUNNYVALE","State":"CA","Zip":"94085","Country":"US"}]`
	err = json.Unmarshal([]byte(rcvd1), &data1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data1)

}
