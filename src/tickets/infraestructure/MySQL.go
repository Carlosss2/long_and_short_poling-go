package infraestructure

import (
	"database/sql"
	"fmt"
	"projectC1/src/tickets/domain"

	
)

type MySQL struct{
	DB *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{DB: db}
}

func (mysql *MySQL) Save(client string, total string) error {
	_, err := mysql.DB.Exec("INSERT INTO tickets (client, total) VALUES (?, ?)", client, total)
	if err != nil {
		return fmt.Errorf("[MySQL] Error al guardar el ticket : %w", err)
	}
	return nil
}

func(mysql *MySQL) GetAll()([]domain.Ticket,error){
	rows,err := mysql.DB.Query("SELECT * FROM tickets")

	if err != nil{
		return nil,err
	}
	defer rows.Close()

	var tickets []domain.Ticket

	for rows.Next(){
		var ticket domain.Ticket
		err := rows.Scan(&ticket.Id,&ticket.Client,&ticket.Total)
		if err != nil{
			return nil,err
		}
		tickets = append(tickets, ticket)

	}
	if err := rows.Err();err != nil{
		return nil,err
	}
	return tickets,nil
}


func (mysql *MySQL)Delete(id int32)(error){
	query := "DELETE FROM tickets WHERE id = ?"
	result,err := mysql.DB.Exec(query,id)
	if err != nil{
		return err
	}
	rowsAffected,err := result.RowsAffected()
	if err != nil{
		return err
	}
	if rowsAffected ==0{
		return fmt.Errorf("No se encontro el ticket con ID:",id)
	}
	fmt.Println("Ticket Elimonado")
	return nil
}

func (mysql *MySQL) Update(id int32,client string,total string) error{
	query := "UPDATE tickets SET client = ?, total =? WHERE id = ? "

	result,err := mysql.DB.Exec(query,client,total,id)
	if err != nil{
		return fmt.Errorf("[MySQL] error alm verificar las filas afectadas:",err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("[MySQL] Error al verificar filas afectadas: %w", err)
	}

	if rowsAffected == 0{
		return fmt.Errorf("No se encontro el Ticket")
	}

	fmt.Println("Ticket actualizado")
	return nil
}