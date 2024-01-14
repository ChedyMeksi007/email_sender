package emailhtml

import( 
	gomail "gopkg.in/gomail.v2"
)

func HtmlEmailCreator(to []string, from string, htmlData string, imagepath string){
    msg := gomail.NewMessage()
    msg.SetHeader("From", from)
    msg.SetHeader("To", to)
    msg.SetHeader("Subject", )
    msg.SetBody("text/html", htmlData)
    msg.Attach(imagepath)	

}
