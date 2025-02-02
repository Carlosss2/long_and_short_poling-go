package application

import "projectC1/src/tickets/domain"

type UpdateTicket struct {
	db domain.ITicket
}

func NewUpdateTicket(db domain.ITicket)*UpdateTicket{
	return &UpdateTicket{db: db}
}
func (uc *UpdateTicket) Execute(id int32, client string,total string) error{
	return uc.db.Update(id,client,total)
}