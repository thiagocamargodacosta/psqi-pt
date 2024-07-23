package exporter

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/thiagocamargodacosta/psqi-pt/reporting"
)

// Strings that will populate the header of the output report
const (
	date_column  string = "Carimbo de data/hora"
	email_column string = "Endereço de e-mail"
	c1_column    string = "Componente 1 - Qualidade subjetiva do sono"
	c2_column    string = "Componente 2 - Latência do sono"
	c3_column    string = "Componente 3 - Duração do sono"
	c4_column    string = "Componente 4 - Eficiência do sono"
	c5_column    string = "Componente 5 - Distúrbios do sono"
	c6_column    string = "Componente 6 - Uso de medicação para dormir"
	c7_column    string = "Componente 7 - Sonolência e disfunção diurnas"
	c8_column    string = "Componente 8 - Qualidade do sono - valor global"
)

var header []string = []string{date_column, email_column, c1_column, c2_column, c3_column, c4_column, c5_column, c6_column, c7_column, c8_column}

func WriteOutputCSV(outputFilename string, reports []reporting.Report) error {

	n := len(reports)
	records := make([][]string, n+1)
	records[0] = header

	for i, r := range reports {
		records[i+1] = []string{r.Date, r.Email, r.Component1, r.Component2, r.Component3, r.Component4, r.Component5, r.Component6, r.Component7, r.GlobalValue}
	}

	file, err := os.Create(outputFilename)

	if err != nil {
		log.Fatalf("unable to create output file", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	w.WriteAll(records)

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}

	return nil
}
