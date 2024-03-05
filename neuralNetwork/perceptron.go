package neuralnetwork

import (
	"math"
)

type Perceptron struct {
	One Matrix

	Two Matrix

	Out Matrix

	check map[string]string
}

func (perceptron *Perceptron) PerceptronGetRan(inputSize int, outputSize int) {

	perceptron.One.RandomGetmat(4, inputSize)

	perceptron.Two.RandomGetmat(4, inputSize+4)

	perceptron.Out.RandomGetmat(outputSize, inputSize+8)

	perceptron.check["000"] = "00"
	perceptron.check["001"] = "01"
	perceptron.check["010"] = "10"
	perceptron.check["011"] = "11"
	perceptron.check["100"] = "00"
	perceptron.check["101"] = "01"
	perceptron.check["110"] = "10"
	perceptron.check["111"] = "11"
}

func Activate(weight Matrix, values Vector) Vector {

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

		if result.Data[i] >= bias.Data[i] {
			result.Data[i] = 0
			result.StringView += "0"
		} else {
			result.Data[i] = 1
			result.StringView += "1"
		}

	}

	return result

}

func (perceptron *Perceptron) PerceptronAction(input Vector) Vector {

	//first layer
	valuesonOne := Activate(perceptron.One, input)

	//second layer
	valuestoTwo := VecConcat(input, valuesonOne)
	valuesonTwo := Activate(perceptron.Two, valuestoTwo)

	//out Layer
	valuestoOut := VecConcat(input, valuesonOne, valuesonTwo)
	valuesonOut := Activate(perceptron.Out, valuestoOut)

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

func (perceptron *Perceptron) PerceptronGet(inputSize int, outputSize int, code []int) {

	borderOne := 4 * inputSize
	perceptron.One.GetMat(4, inputSize, code[0:borderOne])

	borderTwo := 4*(inputSize+4) + borderOne
	perceptron.Two.GetMat(4, inputSize+4, code[borderOne:borderTwo])

	perceptron.Out.GetMat(outputSize, inputSize+8, code[borderTwo:])

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
