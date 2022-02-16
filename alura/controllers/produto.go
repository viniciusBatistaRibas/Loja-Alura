package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"git.hub/vinicius/alura/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	todos_os_produtos := models.Busca_produtos()
	temp.ExecuteTemplate(w, "Index", todos_os_produtos)

}
func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		preco_float, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão de preço: (controllers/produto/Insert) ", err)
		}
		quantidade_int, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade: (controllers/produto/insert)", err)
		}
		models.CriarNovoProduto(nome, descricao, preco_float, quantidade_int)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get(("id"))
	models.DeletaProduto(idProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	produto := models.EditaProduto(id)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func convert(id string) int{
	id_conv, err  := strconv.Atoi(id)
	if err != nil {
		log.Println("Erro na convesão do ID para int: (controllers/produto/update)", err)
	}
	return id_conv
}
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")
		id := r.FormValue("id")


		id_conv:= convert(id)
       
        preco_conv, err := strconv.ParseFloat(preco, 64)
        if err != nil {
            log.Println("Erro na convesão do preço para float64:(controllers/produto/update)", err)
        }

        quantidade_conv, err := strconv.Atoi(quantidade)
        if err != nil {
            log.Println("Erro na convesão da quantidade para int:(controllers/produto/update)", err)
        }
		models.AtualizaProduto(id_conv, nome, descricao, preco_conv, quantidade_conv)

	}
	http.Redirect(w, r, "/", 301)
}
