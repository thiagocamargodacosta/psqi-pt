package scoring

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/thiagocamargodacosta/psqi-pt/form"
)

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

// Sleep quality global value
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

	switch sum {
	case 0:
		score = 0
	case 1:
		score = 1
	case 2:
		score = 1
	case 3:
		score = 2
	case 4:
		score = 2
	case 5:
		score = 3
	case 6:
		score = 3
	}
	return score
}

// Component 3 - Sleep duration
func SleepDuration(entry form.Entry) int {

	PSQI4, _ := strconv.Atoi(entry.PSQI4)

	if PSQI4 > 7 {
		return 0
	} else if PSQI4 >= 6 {
		return 1
	} else if PSQI4 >= 5 {
		return 2
	} else { // entry.PSQI4 < 5
		return 3
	}

}

// Component 4 - Sleep efficiency
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

	if efficiency >= 85.0 {
		return 0
	} else if efficiency >= 75.0 {
		return 1
	} else if efficiency >= 65.0 {
		return 2
	} else { // efficiency < 65
		return 3
	}
}

// Component 5 - Sleep disturbances

func SleepDisturbances(entry form.Entry) int {
	var sum int
	var score int

	switch entry.PSQI5b {
	case "Nenhuma no último mês":
		sum += 0
	case "Menos de 1 vez/ semana":
		sum += 1
	case "1 ou 2 vezes/ semana":
		sum += 2
	case "3 ou mais vezes/ semana":
		sum += 3
	}

	switch entry.PSQI5c {
	case "Nenhuma no último mês":
		sum += 0
	case "Menos de 1 vez/ semana":
		sum += 1
	case "1 ou 2 vezes/ semana":
		sum += 2
	case "3 ou mais vezes/ semana":
		sum += 3
	}

	switch entry.PSQI5d {
	case "Nenhuma no último mês":
		sum += 0
	case "Menos de 1 vez/ semana":
		sum += 1
	case "1 ou 2 vezes/ semana":
		sum += 2
	case "3 ou mais vezes/ semana":
		sum += 3
	}

	switch entry.PSQI5e {
	case "Nenhuma no último mês":
		sum += 0
	case "Menos de 1 vez/ semana":
		sum += 1
	case "1 ou 2 vezes/ semana":
		sum += 2
	case "3 ou mais vezes/ semana":
		sum += 3
	}

	switch entry.PSQI5f {
	case "Nenhuma no último mês":
		sum += 0
	case "Menos de 1 vez/ semana":
		sum += 1
	case "1 ou 2 vezes/ semana":
		sum += 2
	case "3 ou mais vezes/ semana":
		sum += 3
	}

	switch entry.PSQI5g {
	case "Nenhuma no último mês":
		sum += 0
	case "Menos de 1 vez/ semana":
		sum += 1
	case "1 ou 2 vezes/ semana":
		sum += 2
	case "3 ou mais vezes/ semana":
		sum += 3
	}

	switch entry.PSQI5h {
	case "Nenhuma no último mês":
		sum += 0
	case "Menos de 1 vez/ semana":
		sum += 1
	case "1 ou 2 vezes/ semana":
		sum += 2
	case "3 ou mais vezes/ semana":
		sum += 3
	}

	switch entry.PSQI5i {
	case "Nenhuma no último mês":
		sum += 0
	case "Menos de 1 vez/ semana":
		sum += 1
	case "1 ou 2 vezes/ semana":
		sum += 2
	case "3 ou mais vezes/ semana":
		sum += 3
	}

	switch entry.PSQI5j {
	case "Nenhuma no último mês":
		sum += 0
	case "Menos de 1 vez/ semana":
		sum += 1
	case "1 ou 2 vezes/ semana":
		sum += 2
	case "3 ou mais vezes/ semana":
		sum += 3
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
