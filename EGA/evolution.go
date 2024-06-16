package EGA

import (
	nextgeneration "EZ/EGA/nextGeneration"
	startpopulation "EZ/EGA/startPopulation"
	"EZ/graph"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func Evolution(size int, xlfile *excelize.File) int {

	nextGen := startpopulation.CreateStartpopulation(size, 3, 2)

	numberGen := 0

	for i := 1; ; i++ {
		nextGen = nextgeneration.GetnextGen(nextGen, i, xlfile)
		numberGen++
		if nextGen[0].Fitness == 0 {
			fmt.Println(nextGen[0])
			graph.AdjacencyMatrix(nextGen[0].Code.Data)
			break
		}
	}

	return numberGen
}
