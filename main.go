package main 

import(
	"fmt"
	"github.com/ChedyMeksi007/email_sender/excelreader"
)

func main(){
	fmt.Println(excelreader.ExcelReader("test.xlsx"))
}

