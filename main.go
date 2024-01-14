package main

import (
	"fmt"
	"os"
	
	gomail "gopkg.in/gomail.v2"

	"github.com/ChedyMeksi007/email_sender/utils"
	"github.com/joho/godotenv"
)


func init(){

	err := godotenv.Load(".env")

	if err != nil{
		log.Fatal(err)
	}

}
func main() {

	
	//recepients 
	to := excelreader.ExcelReader("test.xlsx")
	
	from := os.Getenv("FROM")
	pwd := os.Getenv("PWD")

	msg := messageCreator(from,to)

}
