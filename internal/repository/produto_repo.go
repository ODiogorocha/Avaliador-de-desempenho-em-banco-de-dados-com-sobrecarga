package repository

import (
	"crud-desempenho-go/internal/model"
	"database/sql"
	"log"
)

type ProdutoRepository struct {
	db *sql.DB
}

func NewProdutoRepository(db *sql.DB) *ProdutoRepository {
	return &ProdutoRepository{db: db}
}

func (r *ProdutoRepository) Criar(p *model.Produto) error {
	query := "INSERT INTO produtos (nome, preco, stock) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, p.Nome, p.Preco, p.Stock)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	p.ID = int(id)
	return nil
}

func (r *ProdutoRepository) BuscarPorID(id int) (*model.Produto, error) {
	var p model.Produto
	query := "SELECT id, nome, preco, stock FROM produtos WHERE id = ?"

	log.Printf("[DEBUG] Tentando buscar no SQLite o ID: %d", id)

	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Nome, &p.Preco, &p.Stock)
	if err == sql.ErrNoRows {
		log.Printf("[DEBUG] ID %d NÃO ENCONTRADO na tabela produtos.", id)
		return nil, nil
	}
	if err != nil {
		log.Printf("[DEBUG] ERRO SQL ao buscar ID %d: %v", id, err)
		return nil, err
	}

	log.Printf("[DEBUG] Produto encontrado com sucesso: ID %d", p.ID)
	return &p, nil
}
