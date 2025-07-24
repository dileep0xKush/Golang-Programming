package condition

import "fmt"

func Switchstatement() {

	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("End of the work week")
	default:
		fmt.Println("Midweek day")
	}

}
