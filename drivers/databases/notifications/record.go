package users

type InvoiceDetail struct {
	Email   string
	Name    string
	Amount  int
	EventID int
}

type Invoice struct {
	Name string
}
