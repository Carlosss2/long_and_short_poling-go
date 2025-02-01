package domain

type Ticket struct{
	Id int32
	Client string
	Total string
}

func NewTicket(client string, total string) *Ticket{
	return &Ticket{Id:0,Client: client,Total: total}
}

func (t *Ticket) GetTotal() string{
	return t.Total
}