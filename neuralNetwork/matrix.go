package neuralnetwork

import "math/rand"

type Matrix struct {
	Data [][]int
	N    int
	M    int
}

func (matrix *Matrix) GetMat(n int, m int, code []int) {
	matrix.N = n
	matrix.M = m

	matrix.Data = make([][]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matrix.Data[i] = append(matrix.Data[i], code[i*m+j])
		}
	}
}

func (matrix *Matrix) RandomGetmat(n int, m int) {

	matrix.N = n
	matrix.M = m

	matrix.Data = make([][]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matrix.Data[i] = append(matrix.Data[i], rand.Intn(2))
		}
	}
}
