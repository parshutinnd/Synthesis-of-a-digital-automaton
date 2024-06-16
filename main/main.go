package main

import (
	"EZ/EGA"
	"fmt"

	"github.com/xuri/excelize/v2"
)

var Automata = new(map[string]string)

func main() {

	xlfile := excelize.NewFile()

	fmt.Println(EGA.Evolution(200, xlfile))

	xlfile.SaveAs("data.xlsx")

}
