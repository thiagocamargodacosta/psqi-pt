package scoring

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/thiagocamargodacosta/psqi-pt/form"
)

// Stores the score for each component of the PSQI
type Score struct {
	Component1  int
	Component2  int
	Component3  int
	Component4  int
	Component5  int
	Component6  int
	Component7  int
	GlobalValue int
}

// Returns the sleep quality score for a given form entry
func SleepQualityScore(entry form.Entry) Score {

	var s Score = Score{
		Component1:  SubjectiveSleepQuality(entry),
		Component2:  SleepLatency(entry),
		Component3:  SleepDuration(entry),
		Component4:  SleepEfficiency(entry),
		Component5:  SleepDisturbances(entry),
		Component6:  SleepMedicationUsage(entry),
		Component7:  SleepnessAndDisfunctions(entry),
		GlobalValue: 0,
	}

	s.GlobalValue = s.Component1 + s.Component2 + s.Component3 + s.Component4 + s.Component5 + s.Component6 + s.Component7

	return s
}

// Component 1 - Subjective sleep quality
// Returns the score for the answer given to the PSQI6 question
func SubjectiveSleepQuality(entry form.Entry) int {
	var score int

	switch entry.PSQI6 {
	case "Muito boa":
		score = 0
	case "Boa":
		score = 1
	case "Ruim":
		score = 2
	case "Muito ruim":
		score = 3
	}

	return score
}

// Component 2 - Sleep latency
// Returns the score for the answers given to the PSQI2 and PSQI5a questions
func SleepLatency(entry form.Entry) int {
	var sum int
	var score int

	PSQI2, _ := strconv.Atoi(entry.PSQI2)

	if PSQI2 <= 15 {
		sum += 0
	} else if PSQI2 <= 30 {
		sum += 1
	} else if PSQI2 <= 60 {
		sum += 2
	} else { // entry.PSQI2 > 60
		sum += 3
	}

	switch entry.PSQI5a {
	case "Nunca":
		sum += 0
	case "Menos 1 vez/ semana":
		sum += 1
	case "1 ou 2 vezes/ semana":
		sum += 2
	case "3 ou mais vezes/ semana":
		sum += 3
	}

	if sum == 0 {
		score = 0
	} else if sum <= 2 {
		score = 1
	} else if sum <= 4 {
		score = 2
	} else {
		score = 3
	}

	return score
}

// Component 3 - Sleep duration
// Returns the score for the answer given to the PSQI4 question
func SleepDuration(entry form.Entry) int {

	var score int
	PSQI4, _ := strconv.Atoi(entry.PSQI4)

	if PSQI4 > 7 {
		score = 0
	} else if PSQI4 >= 6 {
		score = 1
	} else if PSQI4 >= 5 {
		score = 2
	} else { // entry.PSQI4 < 5
		score = 3
	}

	return score
}

// Component 4 - Sleep efficiency
// Returns the score for the answers given to the PSQI4, PSQI3, and PSQI1 questions
func SleepEfficiency(entry form.Entry) int {

	hoursSlept, _ := strconv.Atoi(entry.PSQI4)
	bedTime, _ := strconv.Atoi(entry.PSQI1)
	wakeTime, _ := strconv.Atoi(entry.PSQI3)

	PSQI3 := entry.PSQI3 + "h"
	PSQI1 := entry.PSQI1 + "h"

	// Determine if the person sleeps and wakes on the same day
	var sleepDay int = 1
	var wakeUpDay int = 2

	if bedTime <= wakeTime { // person is sleeping on the same day of waking up
		sleepDay = 2
	}

	wakeUp, err := time.ParseDuration(PSQI3)

	if err != nil {
		log.Println(err)
	}

	lieDown, err := time.ParseDuration(PSQI1)

	if err != nil {
		log.Println(err)
	}

	start := time.Date(2024, time.January, sleepDay, int(lieDown.Hours()), 0, 0, 0, time.UTC)
	end := time.Date(2024, time.January, wakeUpDay, int(wakeUp.Hours()), 0, 0, 0, time.UTC)

	hoursInBed := end.Sub(start)

	var efficiency float64 = float64(hoursSlept) / hoursInBed.Hours() * float64(100)

	fmt.Printf("Sleep Efficiency:\t%s\t%s\tPSQI4 = %d\tPSQI3 = %s\tPSQI1 = %s\tHoursInBed = %s\tEfficiency = %.2f\n", entry.Date, entry.Email, hoursSlept, PSQI3, PSQI1, hoursInBed, efficiency)

	var score int

	if efficiency >= 85.0 {
		score = 0
	} else if efficiency >= 75.0 {
		score = 1
	} else if efficiency >= 65.0 {
		score = 2
	} else { // efficiency < 65
		score = 3
	}

	return score
}

// Component 5 - Sleep disturbances
// Returns the score for the answers given to the PSQI5b, ... , PSQI5j questions
func SleepDisturbances(entry form.Entry) int {
	var sum int
	var score int
	var answers []string = []string{entry.PSQI5b, entry.PSQI5c, entry.PSQI5d, entry.PSQI5e, entry.PSQI5f, entry.PSQI5g, entry.PSQI5h, entry.PSQI5i, entry.PSQI5j}

	for _, answer := range answers {

		switch answer {
		case "Nenhuma no último mês":
			sum += 0
		case "Menos de 1 vez/ semana":
			sum += 1
		case "1 ou 2 vezes/ semana":
			sum += 2
		case "3 ou mais vezes/ semana":
			sum += 3
		}
	}

	if sum == 0 {
		score = 0
	} else if sum < 10 {
		score = 1
	} else if sum <= 18 {
		score = 2
	} else { // sum >= 19 && sum <= 27
		score = 3
	}

	return score
}

// Component 6 - Usage of sleep medication
// Returns the score for the answer given to the PSQI7 question
func SleepMedicationUsage(entry form.Entry) int {
	var score int

	switch entry.PSQI7 {
	case "Nenhuma no último mês":
		score = 0
	case "Menos de 1 vez/ semana":
		score = 1
	case "1 ou 2 vezes/ semana":
		score = 2
	case "3 ou mais vezes/ semana":
		score = 3
	}

	return score
}

// Component 7 - Sleepness and disfunctions
// Returns the score for the answers given to the PSQI8 and PSQI9 questions
func SleepnessAndDisfunctions(entry form.Entry) int {
	var sum int
	var score int

	switch entry.PSQI8 {
	case "Nenhuma no último mês":
		sum = 0
	case "Menos de 1 vez/ semana":
		sum = 1
	case "1 ou 2 vezes/ semana":
		sum = 2
	case "3 ou mais vezes/ semana":
		sum = 3
	}

	switch entry.PSQI9 {
	case "Nenhuma dificuldade":
		sum += 0
	case "Um problema leve":
		sum += 1
	case "Um problema razoável":
		sum += 2
	case "Um grande problema":
		sum += 3
	}

	if sum == 0 {
		score = 0
	} else if sum <= 2 {
		score = 1
	} else if sum <= 4 {
		score = 2
	} else { // sum <= 6
		score = 3
	}

	return score
}
