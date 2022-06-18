package models

import (
	"fmt"
	"store/db"
)

type Products struct {
	Id, Quantidade  int
	Nome, Descricao string
	Preco           float64
}

func GetAllProducts() []Products {
	db := db.ConnectDB()
	dbprodutos, err := db.Query("select * from produtos order by id asc")

	if err != nil {
		panic(err.Error())
	}

	pd := Products{}
	produtos := []Products{}

	for dbprodutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = dbprodutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		pd.Id = id
		pd.Nome = nome
		pd.Descricao = descricao
		pd.Preco = preco
		pd.Quantidade = quantidade

		produtos = append(produtos, pd)
	}
	defer db.Close()
	return produtos
}

func CreateProduct(nome string, descricao string, preco float64, quantidade int) {
	db := db.ConnectDB()
	pre_query, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES($1,$2,$3,$4)")
	if err != nil {
		fmt.Println(err)
	}
	pre_query.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeleteProduct(id_product int) {
	db := db.ConnectDB()

	pre_query, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		fmt.Println(err)
	}
	pre_query.Exec(id_product)
	db.Close()
}

func GetProduct(id_product string) Products {
	db := db.ConnectDB()
	p := Products{}

	product, err := db.Query("select * from produtos where id=$1", id_product)

	if err != nil {
		panic(err.Error())
	}

	for product.Next() {
		var nome, descricao string
		var preco float64
		var id, quantidade int

		err := product.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
	}

	return p
}

func Update(id int, nome string, descricao string, preco float64, quantidade int) {
	db := db.ConnectDB()

	updateProduct, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
