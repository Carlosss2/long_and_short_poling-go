package infraestructure

import (
	"database/sql"
	"fmt"
	"projectC1/src/products/domain"

	
)

type MySQL struct{
	DB *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{DB: db}
}

func (mysql *MySQL) Save(name string, price float32) error {
	_, err := mysql.DB.Exec("INSERT INTO products (name, price) VALUES (?, ?)", name, price)
	if err != nil {
		return fmt.Errorf("[MySQL] Error al guardar el producto: %w", err)
	}
	return nil
}

func (mysql *MySQL) GetAll()([]domain.Product,error) {
	
	rows,err := mysql.DB.Query("SELECT * FROM products")

	if err != nil{
		return nil,err
	}

	defer rows.Close()

	var products []domain.Product

	for rows.Next(){
		var product domain.Product
		err := rows.Scan(&product.Id,&product.Name,&product.Price)
		if err != nil{
			return nil,err
		}
		products = append(products, product)

	}
	if err := rows.Err();err != nil{
		return nil, err
	}
	return products,nil

}

func (mysql *MySQL)GetById(id int32)(domain.Product,error){
	var productById domain.Product

	query := "SELECT id, name, price FROM products WHERE id=?"
	row := mysql.DB.QueryRow(query,id)

	err:= row.Scan(&productById.Id,&productById.Name,&productById.Price)
	if err != nil{
		if err == sql.ErrNoRows{
			return productById, fmt.Errorf("producto con id no encontrado",id)
		}
		return productById,err
	}
	
	return productById,nil
}

func (mysql *MySQL)DeleteProduct(id int32)(error){
	query := "DELETE FROM products WHERE id = ?"
	result,err := mysql.DB.Exec(query,id)
	if err != nil{
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err !=nil{
		return err
	}
	if rowsAffected ==0{
		return fmt.Errorf("No se encontro el producto con ID:",id)
	}
	fmt.Println("Producto eliminado")
	return nil
}

func (mysql *MySQL) UpdateProduct(id int32, name string, price float32) error {
	// Definir la consulta SQL para actualizar un producto por su id
	query := "UPDATE products SET name = ?, price = ? WHERE id = ?"
	
	// Ejecutar la consulta
	result, err := mysql.DB.Exec(query, name, price, id)
	if err != nil {
		return fmt.Errorf("[MySQL] Error al actualizar el producto: %w", err)
	}

	// Verificar si realmente se actualizó alguna fila
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("[MySQL] Error al verificar filas afectadas: %w", err)
	}

	// Si no se encontró ningún producto con el id, devolver un error
	if rowsAffected == 0 {
		return fmt.Errorf("[MySQL] No se encontró un producto con el id %d", id)
	}

	// Confirmar que el producto fue actualizado
	fmt.Println("Producto actualizado con éxito")
	return nil
}