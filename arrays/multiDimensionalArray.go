package arrays

import "fmt"

func MultiDimensionalArray() {

	var matrix [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("matrix[%d][%d] = %d\n", i, j, matrix[i][j])
		}
	}

}
