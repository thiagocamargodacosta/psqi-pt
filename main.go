package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/thiagocamargodacosta/psqi-pt/exporter"
	"github.com/thiagocamargodacosta/psqi-pt/form"
	"github.com/thiagocamargodacosta/psqi-pt/reporting"
	"github.com/thiagocamargodacosta/psqi-pt/scoring"
)

// const filename = `PSQI-BR (respostas) - Respostas ao formulário 1.csv`
const filename = "PSQI-BR  (respostas) - Respostas ao formulário 1.csv"
const outputFilename = `psqi-pt-scoring-results.csv`

func main() {

	// Hardcode filename
	//var filename string = "PSQI-BR  (respostas) - Respostas ao formulário 1.csv"

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// Read the csv file
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	// Convert csv to a list
	forms := form.CreateFormEntry(data)
	reports := make([]reporting.Report, len(forms))

	for i, entry := range forms {
		var s scoring.Score = scoring.SleepQualityScore(entry)
		var r reporting.Report = reporting.Report{
			Date: entry.Date, Email: entry.Email, Component1: strconv.Itoa(s.Component1), Component2: strconv.Itoa(s.Component2), Component3: strconv.Itoa(s.Component3), Component4: strconv.Itoa(s.Component4), Component5: strconv.Itoa(s.Component5), Component6: strconv.Itoa(s.Component6), Component7: strconv.Itoa(s.Component7), GlobalValue: strconv.Itoa(s.GlobalValue)}
		reports[i] = r
	}

	err = exporter.WriteOutputCSV(outputFilename, reports)

	if err != nil {
		log.Fatalf("error while writing output csv", err)
	}

}
