package nextgeneration

import (
	"EZ/EGA/structs"
	"fmt"
	"math/rand"
	"sort"

	"github.com/xuri/excelize/v2"
)

func GetnextGen(previousGen []structs.Individual, numGen int, xlfile *excelize.File) []structs.Individual {

	result := make([]structs.Individual, len(previousGen)/2)
	reproduction := make([]structs.Individual, len(previousGen)/2)

	for i := 0; i < len(previousGen)/4; i += 2 {
		result[i] = previousGen[i]
		result[i+1] = previousGen[i+1]
		reproduction[i], reproduction[i+1] = crossover(previousGen[i], previousGen[i+1])
	}

	for i := len(previousGen) / 4; i < len(previousGen)/2; i += 2 {

		index1 := rand.Intn(len(previousGen)-len(previousGen)/4) + len(previousGen)/4 - 1
		index2 := rand.Intn(len(previousGen)-len(previousGen)/4) + len(previousGen)/4 - 1

		result[i] = previousGen[index1]
		result[i+1] = previousGen[index2]
		reproduction[i], reproduction[i+1] = crossover(previousGen[index1], previousGen[index2])
	}

	reproduction = mutation(reproduction)

	reproduction = selection(reproduction)

	result = append(result, reproduction...)

	sort.Slice(result, func(i, j int) bool { return result[i].Fitness < result[j].Fitness })

	numGen++

	for i := range result {

		cellName, _ := excelize.CoordinatesToCellName(i+5, numGen+2)

		if err := xlfile.SetCellFloat("Sheet1", cellName, result[i].Fitness, 2, 32); err != nil {
			fmt.Println(err)
		}
	}

	return result
}
