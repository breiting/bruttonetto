package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/ryanuber/columnize"
	"github.com/thedevsaddam/gojsonq"
)

const (
	baseURL = "https://onlinerechner.haude.at/haude/brutto-netto-rechner-2013/json"
	guid    = "EA2C7F6C-D727-49C9-A532-C45190635C23"
)

func main() {

	year := 2019
	minSalary := 45000
	maxSalary := 80000
	step := 5000

	output := []string{
		"Brutto/Monat (14x) | Brutto/Jahr | Netto/Jahr | Gesamt/Jahr",
	}

	for i := minSalary; i <= maxSalary; i += step {

		yearSalary := i
		url := fmt.Sprintf("%s/%d/%d/Brutto/jaehrlich/0/0/Angestellter/0/false/0/0/steiermark/keine/0/false/false/true/%s", baseURL, year, yearSalary, guid)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error in request: %v\n", err)
			os.Exit(1)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error in body: %v\n", err)
		}

		dnBruttoMonth := gojsonq.New().JSONString(string(body)).Find("Brutto_DN.Monat")
		dnBrutto := gojsonq.New().JSONString(string(body)).Find("Brutto_DN.Jahr")
		dnNetto := gojsonq.New().JSONString(string(body)).Find("Netto_DN.Jahr")
		dgSum := gojsonq.New().JSONString(string(body)).Find("Summe_DG.Jahr")

		output = append(output, fmt.Sprintf("%7.2f | %7.2f | %7.2f | %7.2f", dnBruttoMonth.(float64), dnBrutto.(float64), dnNetto.(float64), dnBrutto.(float64)+dgSum.(float64)))
	}

	result := columnize.SimpleFormat(output)
	fmt.Println(result)
}
