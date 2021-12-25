package invoices

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"genVoice/business/invoices"
	email "genVoice/helper/mail"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type Datas struct {
	DataInvoice   Invoices
	InvoiceDetail []InvoiceDetail
}

type Invoices struct {
	ID            int `gorm:"primaryKey"`
	Name          string
	UserID        int
	InvoiceDetail []InvoiceDetail `gorm:"foreignKey:EventID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type InvoiceDetail struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Email        string
	Amount       int
	EventID      int
	SignatureKey string
	Link         string
	Status       string `gorm:"default:Belum Dibayar"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func toInvoiceDetailDomain(invoice Datas) invoices.DatasDomain {

	a := invoices.DatasDomain{}
	a.DataInvoice.ID = invoice.DataInvoice.ID
	a.DataInvoice.Name = invoice.DataInvoice.Name
	a.DataInvoice.UserID = invoice.DataInvoice.UserID
	a.DataInvoice.CreatedAt = invoice.DataInvoice.CreatedAt
	a.DataInvoice.UpdatedAt = invoice.DataInvoice.UpdatedAt

	for _, v := range invoice.InvoiceDetail {
		email.Email(v.Email, v.Name, v.Link, strconv.Itoa(v.Amount), invoice.DataInvoice.Name)
		a.InvoiceDetail = append(a.InvoiceDetail, invoices.InvoiceDetailDomain{ID: v.ID, Name: v.Name, Email: v.Email, Amount: v.Amount,
			EventID: v.EventID, Link: v.Link, Status: v.Status, CreatedAt: v.CreatedAt, UpdatedAt: v.UpdatedAt})
	}

	return a
}

func fromInvoiceDomain(domain *invoices.DatasDomain) Invoices {

	result := Invoices{}
	result.ID = domain.DataInvoice.ID
	result.Name = domain.DataInvoice.Name
	result.UserID = domain.DataInvoice.UserID
	result.CreatedAt = domain.DataInvoice.CreatedAt
	result.UpdatedAt = domain.DataInvoice.UpdatedAt

	return result
}

func fromInvoiceDetailDomain(domain *invoices.DatasDomain, id int) []InvoiceDetail {

	resultInvoiceDetail := []InvoiceDetail{}
	for _, e := range domain.InvoiceDetail {
		resultInvoiceDetail = append(resultInvoiceDetail, InvoiceDetail{ID: e.ID, Name: e.Name, Email: e.Email,
			Amount: e.Amount, EventID: id, Link: e.Link, Status: e.Status, CreatedAt: e.CreatedAt, UpdatedAt: e.UpdatedAt})
	}

	return paymentLink(resultInvoiceDetail)
}

func toListDomain(use Invoices, invoic []InvoiceDetail) invoices.DatasDomain {
	dts := Datas{}
	dts.DataInvoice = use

	for i := 0; i < len(invoic); i++ {
		dts.InvoiceDetail = append(dts.InvoiceDetail, invoic[i])
	}
	return toInvoiceDetailDomain(dts)
}

var s snap.Client

func paymentLink(datas []InvoiceDetail) []InvoiceDetail {
	for i := 0; i < len(datas); i++ {
		uuid := uuid.New().String()
		signature_key := (uuid + "200" + strconv.Itoa(datas[i].Amount) + ".00" + "SB-Mid-server-sYHf9k6xSdZJa780ILj-MYXB")
		hasher := sha512.New()
		hasher.Write([]byte(signature_key))
		sha := hex.EncodeToString(hasher.Sum(nil))
		datas[i].SignatureKey = sha
		fmt.Println("S H A : ", sha)

		s.New("SB-Mid-server-sYHf9k6xSdZJa780ILj-MYXB", midtrans.Sandbox)
		req := &snap.Request{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  uuid,
				GrossAmt: int64(datas[i].Amount),
			},

			CreditCard: &snap.CreditCardDetails{
				Secure: true,
			},
			CustomerDetail: &midtrans.CustomerDetails{
				FName: datas[i].Name,
				LName: "",
				Email: datas[i].Email,
				Phone: "081234567890",
			},
		}

		snapResp, _ := s.CreateTransaction(req)
		z := &snapResp.RedirectURL
		datas[i].Link = *z

	}
	return datas
}
