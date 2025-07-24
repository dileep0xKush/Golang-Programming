package logical

import "fmt"

func VotingEligibility() {

	age := 20
	isCitizen := true

	fmt.Println("Can vote:", age >= 18 && isCitizen)

}
