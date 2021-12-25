package emailSucces

import (
	// "fmt"
	"bytes"
	"html/template"

	// "github.com/jung-kurt/gofpdf"
	"log"

	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "PT. Makmur Subur Jaya <emailanda@gmail.com>"
const CONFIG_AUTH_EMAIL = "timduapuluhenam@gmail.com"
const CONFIG_AUTH_PASSWORD = "tim26hore123"

var variable = "Dewi Novita Sari"

func Email(to string, name string, amount string, event string) {
	t, _ := template.ParseFiles("./invoice.html")
	buffer := new(bytes.Buffer)
	t.Execute(buffer, map[string]string{"username": name, "amount": amount, "event": event})

	// fmt.Println(buffer)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", to)
	mailer.SetAddressHeader("Cc", "tralalala@gmail.com", "Tra Lala La")
	mailer.SetHeader("Subject", "Transaksi Sukses")
	mailer.SetBody("text/html", buffer.String())
	// mailer.Attach("./sample.png")

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}
