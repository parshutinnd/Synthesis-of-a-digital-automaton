package nextgeneration

import (
	"EZ/EGA/structs"
	perceptron "EZ/neuralNetwork"
	"math/rand"
)

func twoPointcrossover(vector1, vector2 perceptron.Vector) (perceptron.Vector, perceptron.Vector) {

	point1 := rand.Intn(vector1.Size / 2)
	point2 := rand.Intn(vector1.Size/2) + vector1.Size/2

	var result1, result2 perceptron.Vector

	result1.Size = vector1.Size
	result2.Size = vector2.Size

	result1.Data = make([]int, vector1.Size)
	result2.Data = make([]int, vector2.Size)

	for i := 0; i < point1; i++ {
		result1.Data[i], result2.Data[i] = vector1.Data[i], vector2.Data[i]
	}

	for i := point1; i < point2; i++ {
		result1.Data[i], result2.Data[i] = vector2.Data[i], vector1.Data[i]
	}

	for i := point2; i < vector1.Size; i++ {
		result1.Data[i], result2.Data[i] = vector1.Data[i], vector2.Data[i]
	}

	return result1, result2
}

func crossover(vector1, vector2 structs.Individual) (structs.Individual, structs.Individual) {

	var result1, ressult2 structs.Individual

	result1.Code, ressult2.Code = twoPointcrossover(vector1.Code, vector2.Code)

	var perceptron1, perceptron2 perceptron.Perceptron

	perceptron1.PerceptronGet(result1.Code.Data)
	perceptron2.PerceptronGet(ressult2.Code.Data)

	result1.Fitness = perceptron1.PerceptronError()
	ressult2.Fitness = perceptron2.PerceptronError()

	return result1, ressult2
}
