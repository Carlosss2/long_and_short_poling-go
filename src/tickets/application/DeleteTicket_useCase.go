package application

import "projectC1/src/tickets/domain"

type DeleteTicket struct {
	db domain.ITicket
}

func NewDeleteTicket(db domain.ITicket)*DeleteTicket{
	return &DeleteTicket{db:db}

}

func (uc *DeleteTicket) Execute(id int32)(error){
	return uc.db.Delete(id)
}