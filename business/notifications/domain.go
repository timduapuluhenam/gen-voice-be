package notifications

import "genVoice/business/invoices"

type Service interface {
	GetNotif(status, signature_key string) error
	GetUserBySignature(signature_key string) (invoices.InvoiceDetailDomain, int, error)
}

type Repository interface {
	GetNotif(status, signature_key string) error
	GetUserBySignature(signature_key string) (invoices.InvoiceDetailDomain, int, error)
}
