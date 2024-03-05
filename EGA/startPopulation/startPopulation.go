package startpopulation

import (
	"EZ/EGA/structs"
	"math/rand"
	"sort"
)

func CreateStartpopulation(size int) []structs.Individual {

	result := make([]structs.Individual, size)

	for i := 0; i < size; i++ {

		var code []int

		for i := 0; i < 62; i++ {
			code = append(code, rand.Intn(2))
		}

		result[i].GetIndividual(code)
	}

	sort.Slice(result, func(i, j int) bool { return result[i].Fitness < result[j].Fitness })

	return result
}
