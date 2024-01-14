package emailhtml

import( 
	"log"
	"io/ioutil"
	gomail "gopkg.in/gomail.v2"
)

func readFile(filename string)(string,error){

	content, err := ioutil.ReadFile(filename)
	if err!= nil{
		return "", err
	}

	return string(content), err
}
func HtmlEmailCreator(h map[string][]string,htmlfilepath string,imagepath string)*gomail.Message{
    msg := gomail.NewMessage()
    msg.SetHeaders(h)
    
    body,err :=  readFile(htmlfilepath)
    
    if err!= nil{
	log.Fatal(err)
    }

    msg.SetBody("text/html",body)
    msg.Attach(imagepath)	

    return msg
}
