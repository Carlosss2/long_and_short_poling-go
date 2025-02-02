package dependencies

import (
	"database/sql"
	"fmt"
	"projectC1/src/helpers"

	"projectC1/src/tickets/application"
	"projectC1/src/tickets/infraestructure"
	"projectC1/src/tickets/infraestructure/controllers"
)

var(
	mySQL infraestructure.MySQL
	db *sql.DB
)

func Init() {
	db, err := helpers.ConnectToDB()

	if err!= nil{
		fmt.Println("server error")
		return
	}
	
	


	mySQL = *infraestructure.NewMySQL(db)
	
	
	
}
func CloseDB() {
	if db != nil {
		db.Close()
		fmt.Println("Conexión a la base de datos cerrada.")
	}
}
func GetCreateTicketController()*controllers.CreateTicketController{
	caseCreateTicket := application.NewCreateTicket(&mySQL)
	return controllers.NewCreateTicketController(caseCreateTicket,&mySQL)
}

func GetGetAllTicketController()*controllers.GetAllTicketController{
	caseGetAllTicket := application.NewGetAllTicket(&mySQL)
	return controllers.NewGetAllTicketController(*caseGetAllTicket)
}

func GetDeleteTicketController()*controllers.DeleteTicketController{
	caseDeleteTicket:= application.NewDeleteTicket(&mySQL)
	return controllers.NewDeleteTicketController(caseDeleteTicket)
}

func GetUpdateTicketController()*controllers.UpdateTicketController{
	caseUpdateTicket := application.NewUpdateTicket(&mySQL)
	return controllers.NewUpdateTicketController(caseUpdateTicket)
}

func GetTicketPollingController() *controllers.TicketPollingController {
	return controllers.NewTicketPollingController(&mySQL)
}