package logical

import "fmt"

func LogicalOR() {

	isWeekend := false
	hasHoliday := true
	fmt.Println("Can relax:", isWeekend || hasHoliday)

}
