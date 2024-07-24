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

// filename stores the input csv file containing the answers the the psqi-br form
const filename = "PSQI-BR  (respostas) - Respostas ao formulário 1.csv"

// outputFilename is the name of the csv file that will be created with the score for each entry in the input csv
const outputFilename = `psqi-pt-scoring-results.csv`

func main() {

	// Open input csv
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

	// Convert csv to a slice
	forms := form.CreateFormEntry(data)
	reports := make([]reporting.Report, len(forms))

	// For each entry read from the input table
	for i, entry := range forms {
		// Produce the scoring
		var s scoring.Score = scoring.SleepQualityScore(entry)
		// Create the report to be written to the output csv file
		var r reporting.Report = reporting.Report{
			Date: entry.Date, Email: entry.Email, Component1: strconv.Itoa(s.Component1), Component2: strconv.Itoa(s.Component2), Component3: strconv.Itoa(s.Component3), Component4: strconv.Itoa(s.Component4), Component5: strconv.Itoa(s.Component5), Component6: strconv.Itoa(s.Component6), Component7: strconv.Itoa(s.Component7), GlobalValue: strconv.Itoa(s.GlobalValue)}
		reports[i] = r
	}

	// Write all the scores to the output csv file
	err = exporter.WriteOutputCSV(outputFilename, reports)

	if err != nil {
		log.Fatalf("error while writing output csv", err)
	}

}
