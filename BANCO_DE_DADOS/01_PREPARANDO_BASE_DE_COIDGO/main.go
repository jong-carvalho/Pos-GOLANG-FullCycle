package main

//para exportar esse pacote digitar o comando: go mod init github.com/jong-carvalho/Pos-GOLANG-FullCycle/tree/master/BANCO_DE_DADOS/01_PREPARANDO_BASE_DE_COIDGO
import (
	"fmt"
	"github.com/google/uuid"
)
import "database/sql"

//o underline na frente é porque não vamos utilizar todos os pacotes do mysql, sem ele não roda
import _ "github.com/go-sql-driver/mysql"

//as instruções acima são feitas para criar o arquivo go mod que cria um modulo dizendo a versão do go e também
//guardando todos os imports que vamos utilizar
//após isso devemos dar o comando : go mod tidy
//esse comando atualiza todas as importações

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//INSERINDO NO BANCO
	product := NewProduct("Notebook", 1899.90)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	//UPDATE NO BANCO
	product.Price = 100.0
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	//	SELECT NO BANCO
	p, err := selectProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product: %v, possui o preço de %.2f", p.Name, p.Price)

	// SELECIONANDO TODOS OS PRODUTOS
	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("Product: %v, possui o preço de %.2f", p.Name, p.Price)
	}

	// DELETANDO UM REGISTRO
	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}

}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("INSERT INTO PRODUCTS(ID, NAME, PRICE) VALUES(?, ?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("UPDATE PRODUCTS SET NAME = ?, PRICE = ? WHERE ID = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("SELECT ID, NAME, PRICE FROM PRODUCTS WHERE ID = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	//err = stmt.QueryRowContext(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT ID, NAME, PRICE FROM PRODUCTS")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM PRODUCTS WHERE ID = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
