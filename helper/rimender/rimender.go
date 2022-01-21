package rimender

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"genVoice/drivers/databases/invoices"
	emailRe "genVoice/helper/mail"
	"time"

	cron "github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Time struct {
	Time string
}
type Invoices struct {
	ID            int
	CreatedAt     time.Time
	TimeExpired   int
	ExpiredStatus string
}

type InvoiceDetail struct {
	ID int
	// Name  string
	Email string
}

func Rimender(DB *gorm.DB) {

	// set scheduler berdasarkan zona waktu sesuai kebutuhan
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron.New(cron.WithLocation(jakartaTime))

	// stop scheduler tepat sebelum fungsi berakhir
	defer scheduler.Stop()

	invoice := []invoices.Invoices{}
	invoiceDetail := []invoices.InvoiceDetail{}

	// timeNow:=
	// fmt.Println("Today : ", now.Format("2006-01-02"))

	scheduler.AddFunc("0 07 * * *", func() {
		now := time.Now()
		timeNow := now.Format("2006-01-02")
		DB.Find(&invoice)

		for i := 0; i < len(invoice); i++ {

			// time := invoice[i].CreatedAt.String()
			status := invoice[i].ExpiredStatus
			id := invoice[i].ID

			timeAdd := invoice[i].CreatedAt.AddDate(0, 0, invoice[i].TimeExpired-1).String()
			tenggat := invoice[i].CreatedAt.AddDate(0, 0, invoice[i].TimeExpired).String()
			timeThn := timeAdd[:10]
			fmt.Println(timeThn)

			if timeThn == timeNow && status == "Not Yet" && invoice[i].TimeExpired != 1 {

				DB.Find(&invoiceDetail, "event_id = ?", id)
				for j := 0; j < len(invoiceDetail); j++ {
					email := invoiceDetail[j].Email
					// fmt.Println(email)
					emailRe.Email(email, invoiceDetail[j].Name, invoiceDetail[j].Link, strconv.Itoa(invoiceDetail[j].Amount), invoice[i].Name, tenggat[:10])

				}
				DB.Model(&invoice).Where("id = ?", id).Update("expired_status", "Expired")
			}

		}
	},
	)

	// start scheduler
	scheduler.Start()

	// trap SIGINT untuk trigger shutdown.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

// email.Ema(email)
