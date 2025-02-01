package dependencies

import (
	"database/sql"
	"fmt"
	"projectC1/src/helpers"
	"projectC1/src/products/application"
	"projectC1/src/products/infraestructure"
	"projectC1/src/products/infraestructure/controllers"
)
var(
	mySQL infraestructure.MySQL
	db    *sql.DB
)

func Init(){
	db, err := helpers.ConnectToDB()

	if err != nil{
		fmt.Println("Server error")
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


func GetCreateProductController()*controllers.CreateProductController{
	caseCreateProduct := application.NewCreateProduct(&mySQL)
	return controllers.NewCreateProductController(caseCreateProduct)
}

func GetAllProductController()*controllers.GetAllProductController{
	caseGetAllProduct := application.NewGetAllProduct(&mySQL)
	return controllers.NewGetAllProductController(*caseGetAllProduct)
}


func GetDeleteController()*controllers.DeleteProductController{
	caseDelete := application.NewDeleteProduct(&mySQL)
	return controllers.NewDeleteProductController(caseDelete)
}

func GetUpdateController()*controllers.UpdateProductController{
	caseUpdate := application.NewUpdateProduct(&mySQL)
	return controllers.NewUpdateProductController(caseUpdate)
}