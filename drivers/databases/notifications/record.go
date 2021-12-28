package users

type InvoiceDetail struct {
	Email   string
	Name    string
	Amount  int
	EventID string
}

type Invoice struct {
	Name string
}
