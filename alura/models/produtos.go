package models

import "git.hub/vinicius/alura/db"

type Produto struct {
	Nome, Descrição string
	Preco           float64
	Id, Quantidade  int
}

func Busca_produtos() []Produto {
	db := db.Conecta()

	selectDeProdutos, err := db.Query("select * from itens order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeProdutos.Next() {
		var quantidade int
		var nome, descricao string
		var preco float64
		var id int

		err = selectDeProdutos.Scan(&nome, &descricao, &preco, &quantidade, &id)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descrição = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}
func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.Conecta()
	insereDados, err := db.Prepare("insert into itens (nome, descricao,preco,quantidade) values ($1,$2,$3,$4)")

	if err != nil {
		panic(err.Error())
	}
	insereDados.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.Conecta()
	deletarOProduto, err := db.Prepare("delete from itens where id = $1")
	if err != nil {
		panic(err.Error())
	}
	deletarOProduto.Exec(id)
	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := db.Conecta()
	produtoBanco, err := db.Query("select * from itens where id = $1", id)
	if err != nil {
		panic(err.Error())
	}
	produtoAtt := Produto{}

	for produtoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoBanco.Scan(&nome, &descricao, &preco, &quantidade, &id)
		if err != nil {
			panic(err.Error())
		}
		produtoAtt.Id = id
		produtoAtt.Nome = nome
		produtoAtt.Descrição = descricao
		produtoAtt.Preco = preco
		produtoAtt.Quantidade = quantidade

	}
	defer db.Close()
	return produtoAtt
}


func AtualizaProduto(id int, nome , descricao string, preco float64, quantidade int){
	db := db.Conecta()
	
	AtualizaProduto, err := db.Prepare("update itens set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id =$5")
	if err!=nil{
		panic(err.Error())
	}
	AtualizaProduto.Exec(nome,descricao,preco,quantidade,id)
	defer db.Close()
}