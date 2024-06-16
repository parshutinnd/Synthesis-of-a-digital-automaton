package nextgeneration

import (
	"EZ/EGA/structs"
	neuralnetwork "EZ/neuralNetwork"
	"math/rand"
)

func mutation(generation []structs.Individual) []structs.Individual {

	result := generation

	for i := range generation {

		for j := range result[i].Code.Data {
			a := rand.Intn(100)
			if a < 10 {
				if result[i].Code.Data[j] == 1 {
					result[i].Code.Data[j] = 0
				} else {
					result[i].Code.Data[j] = 1
				}
			}
		}

		var perceptron neuralnetwork.Perceptron

		perceptron.PerceptronGet(result[i].Code.Data)

		result[i].Fitness = perceptron.PerceptronError()
	}

	return result
}
