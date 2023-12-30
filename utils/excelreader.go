package excelreader

import (
	"log"

	"github.com/xuri/excelize/v2"
)

// retuns slice of string from excel file
func ExcelReader(filepath string) []string {

	f, err := excelize.OpenFile(filepath)

	if err != nil {
		log.Fatal(err)
	}

	cols, err := f.GetCols("Tabellenblatt1")

	if err != nil {
		log.Fatal(err)
	}

	var recepients []string

	for _, col := range cols {
		for _, rowCell := range col {
			recepients = append(recepients, rowCell)
		}
	}

	return recepients
}
