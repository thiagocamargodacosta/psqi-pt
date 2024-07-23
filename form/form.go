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

	var form []Entry

	for i, line := range data {

		if i > 0 { // Discard header row

			var rec Entry

			for j, field := range line {

				switch j {
				case 0:
					rec.Date = field
				case 1:
					rec.Email = field
				case 2:
					rec.PSQI1 = field
				case 3:
					rec.PSQI2 = field
				case 4:
					rec.PSQI3 = field
				case 5:
					rec.PSQI4 = field
				case 6:
					rec.PSQI5a = field
				case 7:
					rec.PSQI5b = field
				case 8:
					rec.PSQI5c = field
				case 9:
					rec.PSQI5d = field
				case 10:
					rec.PSQI5e = field
				case 11:
					rec.PSQI5f = field
				case 12:
					rec.PSQI5g = field
				case 13:
					rec.PSQI5h = field
				case 14:
					rec.PSQI5i = field
				case 15:
					rec.PSQI5j = field
				case 16:
					rec.Reasons = field
				case 17:
					rec.PSQI6 = field
				case 18:
					rec.PSQI7 = field
				case 19:
					rec.PSQI8 = field
				case 20:
					rec.PSQI9 = field
				}
			}
			form = append(form, rec)
		}
	}
	return form
}
