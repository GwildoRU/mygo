package main

import (
	"fmt"
	"github.com/plandem/xlsx"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir, fn := filepath.Split(os.Args[1])
	xl, err := xlsx.Open(fn)
	if err != nil {
		log.Fatal(err)
	}

	nfn := fn

	defer func() {
		os.Rename(filepath.Join(dir, fn), filepath.Join(dir, nfn))
	}()

	defer xl.Close()

	sheet := xl.Sheet(0)

	var csvw *os.File
	_, totalRows := sheet.Dimension()

	for rIdx := 6; rIdx < totalRows-1; rIdx++ {
		row := sheet.Row(rIdx).Values()

		if nfn == fn {
			nfn = row[1][6:] + row[1][3:5] + filepath.Ext(fn)
			csvw, err = os.Create(filepath.Join(dir, strings.ReplaceAll(nfn, filepath.Ext(nfn), ".csv")))
			if err != nil {
				log.Fatal(err)
			}
			defer csvw.Close()
		}

		fmt.Fprint(csvw)
		fmt.Fprint(csvw, row[1], "\t")                           //DATA_PL
		fmt.Fprint(csvw, row[10], "\t")                          //SUMM_PL
		fmt.Fprint(csvw, row[11], "\t")                          //PENALTY
		fmt.Fprint(csvw, row[2], "\t")                           //LS
		fmt.Fprint(csvw, row[1][3:5], row[1][8:], "\t")          //PERIOD
		fmt.Fprint(csvw, "03", "\t")                             //COD_RKC
		fmt.Fprint(csvw, row[8], "-", row[0], "\t")              //PD_NUM
		fmt.Fprint(csvw, "-", "\t")                              //INN
		fmt.Fprint(csvw, "-", "\t")                              //KPP
		fmt.Fprint(csvw, row[3], " ", row[4], "\t")              //PLAT
		fmt.Fprint(csvw, row[7], " ", row[8], " ", row[9], "\t") //PRIM
		fmt.Fprint(csvw, nfn, "\t\n")                               //filename
	}
}
