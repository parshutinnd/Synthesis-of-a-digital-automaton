package neuralnetwork

import "math/rand"

type Vector struct {
	Data       []int
	StringView string
	Size       int
}

func (vector *Vector) RandomGetvec(size int) {

	for i := 0; i < size; i++ {
		vector.Data = append(vector.Data, rand.Intn(2))
	}

	vector.Size = size
}

func (vector *Vector) VectorGet(input string) {

	vector.Size = len(input)
	vector.Data = make([]int, vector.Size)
	vector.StringView = input

	for i := range vector.Data {
		vector.Data[i] = int([]rune(input)[i]) - 48
	}
}

func VecConcat(vectors ...Vector) Vector {

	var result Vector

	for _, vector := range vectors {
		result.Data = append(result.Data, vector.Data...)
		result.Size += vector.Size
		result.StringView += vector.StringView
	}

	return result
}
