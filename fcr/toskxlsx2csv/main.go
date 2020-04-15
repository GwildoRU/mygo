package main

import (
	"fmt"
	"github.com/plandem/xlsx"
	"log"
	"os"
	"path/filepath"
)

func main() {
	fn := filepath.Base(os.Args[1])
	xl, err := xlsx.Open(fn)
	if err != nil {
		log.Fatal(err)
	}

	defer xl.Close()

	sheet := xl.Sheet(0)

	_, totalRows := sheet.Dimension()
	for rIdx := 6; rIdx < totalRows-1; rIdx++ {
		row := sheet.Row(rIdx).Values()

		fmt.Print(row[1], "\t")                           //DATA_PL
		fmt.Print(row[1], "\t")                           //DATA_PL
		fmt.Print(row[10], "\t")                          //SUMM_PL
		fmt.Print(row[11], "\t")                          //PENALTY
		fmt.Print(row[2], "\t")                           //LS
		fmt.Print(row[1][3:5], row[1][8:], "\t")          //PERIOD
		fmt.Print("03", "\t")                             //COD_RKC
		fmt.Print(row[8], "-", row[0], "\t")              //PD_NUM
		fmt.Print("-", "\t")                              //INN
		fmt.Print("-", "\t")                              //KPP
		fmt.Print(row[3], " ", row[4], "\t")              //PLAT
		fmt.Print(row[7], " ", row[8], " ", row[9], "\t") //PRIM
		fmt.Print(fn, "\t") //filename

		fmt.Println()
	}

}
