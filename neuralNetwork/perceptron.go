package neuralnetwork

import (
	"math"
)

type Perceptron struct {
	One Matrix

	Two Matrix

	Out Matrix

	layerOne []int

	layerTwo []int

	layerThree []int

	check map[string]string
}

func Activate(weight Matrix, values Vector, neurons []int) Vector {

	var bias Vector
	bias.Size = weight.N
	bias.Data = make([]int, weight.N)

	var result Vector
	result.Size = weight.N
	result.Data = make([]int, weight.N)

	for i := 0; i < weight.N; i++ {
		for j := 0; j < weight.M; j++ {
			bias.Data[i] += weight.Data[i][j]
		}
	}

	for i := 0; i < weight.N; i++ {
		for j := 0; j < weight.M; j++ {
			result.Data[i] += values.Data[j] * weight.Data[i][j]
		}
	}

	for i := range bias.Data {

		if neurons[i] == 1 {
			if result.Data[i] >= bias.Data[i] {
				result.Data[i] = 0
				result.StringView += "0"
			} else {
				result.Data[i] = 1
				result.StringView += "1"
			}
		} else {
			if result.Data[i]%2 == 1 {
				result.Data[i] = 0
				result.StringView += "0"
			} else {
				result.Data[i] = 1
				result.StringView += "1"
			}
		}

	}

	return result

}

func (perceptron *Perceptron) PerceptronAction(input Vector) Vector {

	//first layer
	valuesonOne := Activate(perceptron.One, input, perceptron.layerOne)

	//second layer
	valuestoTwo := VecConcat(input, valuesonOne)
	valuesonTwo := Activate(perceptron.Two, valuestoTwo, perceptron.layerTwo)

	//out Layer
	valuestoOut := VecConcat(input, valuesonOne, valuesonTwo)
	valuesonOut := Activate(perceptron.Out, valuestoOut, perceptron.layerThree)

	return valuesonOut
}

func (perceptron *Perceptron) PerceptronError() float64 {

	var err float64

	for key, value := range perceptron.check {

		var input, output Vector
		var difference float64

		input.VectorGet(key)
		output = perceptron.PerceptronAction(input)

		for i := range output.StringView {
			difference += math.Abs(float64(rune(value[i]) - rune(output.StringView[i])))
		}

		difference = math.Pow(difference, 2)
		err += difference / 2
	}

	return err
}

func (perceptron *Perceptron) PerceptronGet(code []int) {

	var inputSize = 3
	var outputSize = 2

	var weightsBorder = 8 + outputSize

	borderOne := 4*inputSize + weightsBorder
	perceptron.One.GetMat(4, inputSize, code[weightsBorder:borderOne])

	borderTwo := 4*(inputSize+4) + borderOne
	perceptron.Two.GetMat(4, inputSize+4, code[borderOne:borderTwo])

	perceptron.Out.GetMat(outputSize, inputSize+8, code[borderTwo:])

	perceptron.layerOne = code[:4]
	perceptron.layerTwo = code[4:8]
	perceptron.layerThree = code[8:weightsBorder]

	perceptron.check = make(map[string]string)

	perceptron.check["000"] = "00"
	perceptron.check["001"] = "01"
	perceptron.check["010"] = "01"
	perceptron.check["011"] = "10"
	perceptron.check["100"] = "01"
	perceptron.check["101"] = "10"
	perceptron.check["110"] = "10"
	perceptron.check["111"] = "11"

}
