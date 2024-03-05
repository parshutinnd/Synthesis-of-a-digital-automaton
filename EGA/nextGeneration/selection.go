package nextgeneration

import (
	"EZ/EGA/structs"
	"math/rand"
)

func selection(generation []structs.Individual) []structs.Individual {

	var result []structs.Individual

	for i := 0; i < len(generation); i++ {

		min := 100.0
		var best structs.Individual

		for j := 0; j < 3; j++ {
			index := rand.Intn(len(generation))
			if generation[index].Fitness < min {
				best = generation[index]
				min = generation[index].Fitness
			}
		}

		result = append(result, best)
	}

	return result
}
