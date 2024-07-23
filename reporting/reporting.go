package reporting

import "fmt"

// Report will represent the scoring result for a given form
type Report struct {
	Date        string // contains the date from the form submission entry
	Email       string // user's email address for identification purposes
	Component1  string // subjective sleep quality
	Component2  string // sleep latency
	Component3  string // sleep duration
	Component4  string // sleep efficiency
	Component5  string // sleep disturbances
	Component6  string // sleep medication usage
	Component7  string // sleepness and daytime dysfunction
	GlobalValue string // sleep quality
}

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

var report_header [][]string = [][]string{{date_column, email_column, c1_column, c2_column, c3_column, c4_column, c5_column, c6_column, c7_column, c8_column}}

func Print(r Report) {

	fmt.Printf("%s:\t\t\t\t%s\n", date_column, r.Date)
	fmt.Printf("%s:\t\t%s\n", c1_column, r.Component1)
	fmt.Printf("%s:\t\t\t%s\n", c2_column, r.Component2)
	fmt.Printf("%s:\t\t\t\t%s\n", c3_column, r.Component3)
	fmt.Printf("%s:\t\t\t%s\n", c4_column, r.Component4)
	fmt.Printf("%s:\t\t\t%s\n", c5_column, r.Component5)
	fmt.Printf("%s:\t\t%s\n", c6_column, r.Component6)
	fmt.Printf("%s:\t\t%s\n", c7_column, r.Component7)
	fmt.Printf("%s:\t%s\n", c8_column, r.GlobalValue)

}
