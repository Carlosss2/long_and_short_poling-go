package application

import "projectC1/src/tickets/domain"

type GetAllTicket struct {
	db domain.ITicket
}

func NewGetAllTicket(db domain.ITicket)*GetAllTicket{
	return &GetAllTicket{db: db}
}

func (getTicket *GetAllTicket) Execute()([]domain.Ticket,error){
	return getTicket.db.GetAll()
}