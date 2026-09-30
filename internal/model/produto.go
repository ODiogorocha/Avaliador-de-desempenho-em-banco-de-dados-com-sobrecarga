// Produto representa as entidades de dominio
package model

type Produto struct {
	ID    int     `json:"id"`
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
	Stock int     `json:"stock"`
}
