package application

import "projectC1/src/tickets/domain"

type CreateTicket struct {
	db domain.ITicket
}

func NewCreateTicket(db domain.ITicket) *CreateTicket{
	return &CreateTicket{db:db}
}

func (uc *CreateTicket) Execute(client string, total string) error{
	return uc.db.Save(client,total)
}