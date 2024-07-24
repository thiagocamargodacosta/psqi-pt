package form

// Definition of the Entry type related to the form that was built for the study
type Entry struct {
	Date    string
	Email   string
	PSQI1   string
	PSQI2   string
	PSQI3   string
	PSQI4   string
	PSQI5a  string
	PSQI5b  string
	PSQI5c  string
	PSQI5d  string
	PSQI5e  string
	PSQI5f  string
	PSQI5g  string
	PSQI5h  string
	PSQI5i  string
	PSQI5j  string
	Reasons string
	PSQI6   string
	PSQI7   string
	PSQI8   string
	PSQI9   string
}

// Extract each entry from the form responses table and return it as a slice
func CreateFormEntry(data [][]string) []Entry {

	var forms []Entry

	for i, line := range data {

		if i > 0 { // Discard header row

			var e Entry

			for j, field := range line {

				switch j {
				case 0:
					e.Date = field
				case 1:
					e.Email = field
				case 2:
					e.PSQI1 = field
				case 3:
					e.PSQI2 = field
				case 4:
					e.PSQI3 = field
				case 5:
					e.PSQI4 = field
				case 6:
					e.PSQI5a = field
				case 7:
					e.PSQI5b = field
				case 8:
					e.PSQI5c = field
				case 9:
					e.PSQI5d = field
				case 10:
					e.PSQI5e = field
				case 11:
					e.PSQI5f = field
				case 12:
					e.PSQI5g = field
				case 13:
					e.PSQI5h = field
				case 14:
					e.PSQI5i = field
				case 15:
					e.PSQI5j = field
				case 16:
					e.Reasons = field
				case 17:
					e.PSQI6 = field
				case 18:
					e.PSQI7 = field
				case 19:
					e.PSQI8 = field
				case 20:
					e.PSQI9 = field
				}
			}
			forms = append(forms, e)
		}
	}
	return forms
}
