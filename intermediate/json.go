package main

import (
	"encoding/json"
	"fmt"
)

type title struct {
	English string
	French  string
}

type titles []title

func main() {
	jsonData := []byte(`[
		{"English":"Mister", "French": "Monsieur"},
		{"English":"Docter", "French": "Docteur"},
		{"English":"Professer", "French": "Professeur"}
	]`)
	var titles titles
	_ = json.Unmarshal(jsonData, &titles)
	for _, t := range titles {
		fmt.Printf("%s -> %s \n", t.English, t.French)
	}
}
