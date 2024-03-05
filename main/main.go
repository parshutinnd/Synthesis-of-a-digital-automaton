package main

import (
	"EZ/EGA"
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {

	xlfile := excelize.NewFile()

	fmt.Println(EGA.Evolution(200, xlfile))

	xlfile.SaveAs("data.xlsx")

}
