package startpopulation

import (
	"EZ/EGA/structs"
	"math/rand"
	"sort"
)

func CreateStartpopulation(size, inputSize, outputSize int) []structs.Individual {

	result := make([]structs.Individual, size)

	for i := 0; i < size; i++ {

		var code []int

		for i := 0; i < 8+outputSize+outputSize*(inputSize+8)+4*(inputSize+4)+4*inputSize; i++ {
			code = append(code, rand.Intn(2))
		}

		result[i].GetIndividual(code)
	}

	sort.Slice(result, func(i, j int) bool { return result[i].Fitness < result[j].Fitness })

	return result
}
