package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"crud-desempenho-go/internal/model"
	"crud-desempenho-go/internal/repository"
)

type ProdutoHandler struct {
	repo *repository.ProdutoRepository
}

func NewProdutoHandler(repo *repository.ProdutoRepository) *ProdutoHandler {
	return &ProdutoHandler{repo: repo}
}

func (h *ProdutoHandler) CriarProduto(w http.ResponseWriter, r *http.Request) {
	var p model.Produto
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Payload inválido", http.StatusBadRequest)
		return
	}

	if err := h.repo.Criar(&p); err != nil {
		http.Error(w, "Erro interno ao salvar produto", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func (h *ProdutoHandler) ObterProduto(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	p, err := h.repo.BuscarPorID(id)
	if err != nil {
		http.Error(w, "Erro interno ao buscar produto", http.StatusInternalServerError)
		return
	}
	if p == nil {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
