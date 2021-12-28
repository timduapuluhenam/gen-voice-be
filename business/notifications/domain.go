package notifications

type Service interface {
	GetNotif(status, signature_key string) error
}

type Repository interface {
	GetNotif(status, signature_key string) error
}
