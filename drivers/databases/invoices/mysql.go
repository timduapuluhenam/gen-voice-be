package invoices

import (
	"genVoice/business/invoices"

	"gorm.io/gorm"
)

type MysqlInvoiceRepository struct {
	Conn *gorm.DB
}

func NewMysqlInvoiceRepository(conn *gorm.DB) invoices.Repository {
	return &MysqlInvoiceRepository{
		Conn: conn,
	}
}

func (rep *MysqlInvoiceRepository) CreateInvoiceDetail(invoiceDetailDomain *invoices.DatasDomain) (invoices.DatasDomain, error) {

	invoic := fromInvoiceDomain(invoiceDetailDomain)
	// fmt.Print(invoic)
	resultInvoice := rep.Conn.Create(&invoic)

	id := invoic.ID
	invoiceDetail := fromInvoiceDetailDomain(invoiceDetailDomain, id)
	// fmt.Print(invoiceDetail)
	resultInvoiceDetail := rep.Conn.Create(&invoiceDetail)

	if resultInvoice.Error != nil {
		return invoices.DatasDomain{}, resultInvoice.Error
	}
	if resultInvoiceDetail.Error != nil {
		return invoices.DatasDomain{}, resultInvoiceDetail.Error
	}
	// fmt.Print("mysql go invoice nih", invoic, invoiceDetail)
	return toListDomain(invoic, invoiceDetail), nil

}

func (rep *MysqlInvoiceRepository) GetAllByUserID(userID int) ([]invoices.InvoiceDetailDomain, error) {
	invoice := []Invoices{}
	invoiceDetail := []InvoiceDetail{}
	resultInvoice := rep.Conn.Find(&invoice, "user_iD = ?", userID)
	result := []invoices.InvoiceDetailDomain{}
	if resultInvoice.Error != nil {
		return []invoices.InvoiceDetailDomain{}, resultInvoice.Error
	}
	for i := range invoice {
		resultInvoiceDetail := rep.Conn.Find(&invoiceDetail, "event_id = ?", invoice[i].ID)
		if resultInvoiceDetail.Error != nil {
			return []invoices.InvoiceDetailDomain{}, resultInvoiceDetail.Error
		}
		for j := range invoiceDetail {
			result = append(result, invoices.InvoiceDetailDomain{
				InvoiceName:  invoice[i].Name,
				ID:           invoiceDetail[j].ID,
				Name:         invoiceDetail[j].Name,
				SignatureKey: invoiceDetail[j].SignatureKey,
				Email:        invoiceDetail[j].Email,
				Status:       invoiceDetail[j].Status,
				EventID:      invoiceDetail[j].EventID,
				Amount:       invoiceDetail[j].Amount,
				Link:         invoiceDetail[j].Link,
				CreatedAt:    invoiceDetail[j].CreatedAt,
				UpdatedAt:    invoiceDetail[j].UpdatedAt,
			},
			)
		}
	}
	return result, nil
}

func (rep *MysqlInvoiceRepository) DeleteInvoice(invoiceID int) (invoices.InvoiceDomain, error) {

	// invoic := fromInvoiceDomain(invoiceDetailDomain)
	invoice := Invoices{}
	invoiceDomain := invoices.InvoiceDomain{}
	rep.Conn.Find(&invoice, "ID=?", invoiceID)

	invoiceDomain.ID = invoice.ID
	invoiceDomain.UserID = invoice.ID
	invoiceDomain.Name = invoice.Name
	invoiceDomain.TimeExpired = invoice.TimeExpired
	invoiceDomain.CreatedAt = invoice.CreatedAt
	invoiceDomain.UpdatedAt = invoice.UpdatedAt

	// fmt.Println(invoiceDomain)
	rep.Conn.Where("ID=?", invoiceID).Delete(&invoice)

	return invoiceDomain, nil

}


func (rep *MysqlInvoiceRepository) GetInvoiceDetailByID(id string) (invoices.InvoiceDetailDomain, error) {

	invoiceDetail := InvoiceDetail{}
	invoice := Invoices{}
	invoiceDetailDomain := invoices.InvoiceDetailDomain{}
	rep.Conn.First(&invoiceDetail, "id = ?", id)
	rep.Conn.First(&invoice, invoiceDetail.EventID)

	invoiceDetailDomain.InvoiceName = invoice.Name
	invoiceDetailDomain.ID = invoiceDetail.ID
	invoiceDetailDomain.Name = invoiceDetail.Name
	invoiceDetailDomain.Email = invoiceDetail.Email
	invoiceDetailDomain.Amount = invoiceDetail.Amount
	invoiceDetailDomain.EventID = invoiceDetail.EventID
	invoiceDetailDomain.SignatureKey = invoiceDetail.SignatureKey
	invoiceDetailDomain.Link = invoiceDetail.Link
	invoiceDetailDomain.Status = invoiceDetail.Status
	invoiceDetailDomain.CreatedAt = invoiceDetail.CreatedAt
	invoiceDetailDomain.UpdatedAt = invoiceDetail.UpdatedAt

	return invoiceDetailDomain, nil
}

func (rep *MysqlInvoiceRepository) GetAllEventByUserID(userID int) ([]invoices.InvoiceDomain,error) {
	invoice := []Invoices{}
	resultInvoice := rep.Conn.Find(&invoice, "user_iD = ?", userID)
	if resultInvoice.Error != nil {
		return nil, resultInvoice.Error
	}
	return toListInvoiceDomain(invoice),nil
}

