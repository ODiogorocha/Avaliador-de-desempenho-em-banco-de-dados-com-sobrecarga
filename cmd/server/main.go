package main

import (
	"log"
	"net/http"

	"crud-desempenho-go/internal/config"
	"crud-desempenho-go/internal/handler"
	"crud-desempenho-go/internal/repository"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Falha crítica ao iniciar BD: %v", err)
	}
	defer db.Close()

	produtoRepo := repository.NewProdutoRepository(db)
	produtoHandler := handler.NewProdutoHandler(produtoRepo)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /produtos", produtoHandler.CriarProduto)
	mux.HandleFunc("GET /produtos/{id}", produtoHandler.ObterProduto)

	log.Println("Servidor web iniciado na porta :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Erro no servidor HTTP: %v", err)
	}
}
