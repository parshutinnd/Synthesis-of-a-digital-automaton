package graph

import "fmt"

func AdjacencyMatrix(code []int) {
	var result [13][13]int

	var count = 10
	for i := 3; i <= 6; i++ {
		for j := 0; j < 3; j++ {
			result[i][j] = code[count]
			count++
		}
	}

	for i := 7; i <= 10; i++ {
		for j := 0; j < 7; j++ {
			result[i][j] = code[count]
			count++
		}
	}

	for i := 11; i <= 12; i++ {
		for j := 0; j < 11; j++ {
			result[i][j] = code[count]
			count++
		}
	}

	for i := 0; i < 13; i++ {
		for j := 0; j < 13; j++ {
			result[i][j] = result[j][i]
		}
	}

	for i := 0; i < 13; i++ {
		for j := 0; j < 13; j++ {
			fmt.Print(result[i][j])
			fmt.Print(",")
		}
		fmt.Println()
	}
}
