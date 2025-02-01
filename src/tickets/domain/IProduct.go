package domain

type ITicket interface{
	Save(client string, total string) error
	GetAll()([]Ticket,error)
	Delete(id int32)error
	Update(id int32,client string,total string)error
}