package main

import (
	"os"
	"log"

	gomail "gopkg.in/gomail.v2"
	"github.com/ChedyMeksi007/email_sender/excelreader"
	htmlem "github.com/ChedyMeksi007/email_sender/htmlemailcreator"
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
	
	from := make([]string,1,1)
	from = append(from,os.Getenv("FROM"))
	h := make(map[string][]string)

	h["to"]= to
	h["from"] = from
	
	pwd := os.Getenv("PWD")

	msg := htmlem.HtmlEmailCreator(h,"test.html","image.png")
	


	d := gomail.NewDialer("smtp.example.com", 587, from[0], pwd)

	if err:= d.DialAndSend(msg); err!=nil{
		log.Fatal(err)
	}
}
