package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"store/models"
	"strconv"
)

var tpls = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosProdutos := models.GetAllProducts()
	tpls.ExecuteTemplate(w, "Index", todosProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	tpls.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConv, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			panic(err.Error())
		}

		quantidadeConv, err := strconv.Atoi(quantidade)
		if err != nil {
			panic(err.Error())
		}

		models.CreateProduct(nome, descricao, precoConv, quantidadeConv)

	}

	http.Redirect(w, r, "Index", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id_product, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		fmt.Println(err)
	}

	models.DeleteProduct(id_product)
	http.Redirect(w, r, "Index", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	product := models.GetProduct(r.URL.Query().Get("id"))
	tpls.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConv, err := strconv.Atoi(id)
		if err != nil {
			panic(err.Error())
		}
		precoConv, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			panic(err.Error())
		}
		quantidadeConv, err := strconv.Atoi(quantidade)
		if err != nil {
			panic(err.Error())
		}

		models.Update(idConv, nome, descricao, precoConv, quantidadeConv)
		http.Redirect(w, r, "Index", 301)
	}
}
