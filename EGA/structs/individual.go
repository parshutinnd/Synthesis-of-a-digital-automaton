package structs

import (
	perceptron "EZ/neuralNetwork"
)

type Individual struct {
	Code    perceptron.Vector
	Fitness float64
}

func (individual *Individual) GetIndividual(code []int) {

	individual.Code.Data = code
	individual.Code.Size = len(code)

	var perceptron perceptron.Perceptron

	perceptron.PerceptronGet(3, 2, code)

	individual.Fitness = perceptron.PerceptronError()

}
