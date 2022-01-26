package email

import (
	// "fmt"
	"bytes"
	"html/template"

	// "github.com/jung-kurt/gofpdf"
	"log"

	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 25
const CONFIG_SENDER_NAME = "PT. Genvoice Indonesia <emailanda@gmail.com>"
const CONFIG_AUTH_EMAIL = "timduapuluhenam@gmail.com"
const CONFIG_AUTH_PASSWORD = "tim26hore123"

// var variable = "Dewi Novita Sari"

func Email(to string, name string, link string, amount string, event string, tanggal string, invoiceId string) {
	t, err := template.ParseFiles("template3.html")
	if err != nil {
		log.Fatal(err.Error())
	}
	buffer := new(bytes.Buffer)
	t.Execute(buffer, map[string]string{"username": name, "link": link, "amount": amount, "event": event, "tanggal": tanggal, "invoiceId": invoiceId})

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", to)
	mailer.SetAddressHeader("Cc", "tralalala@gmail.com", "Tra Lala La")
	mailer.SetHeader("Subject", "Tagihan")
	mailer.SetBody("text/html", buffer.String())
	// mailer.Attach("./sample.png")

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err = dialer.DialAndSend(mailer)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}
