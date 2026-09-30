package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./trabalho.db")
	if err != nil {
		log.Fatalf("Erro ao abrir banco: %v", err)
	}
	defer db.Close()

	// DROP TABLE garante que a tabela é recriada e o auto-incremento zera (começa do 1)
	_, err = db.Exec(`DROP TABLE IF EXISTS produtos;`)
	if err != nil {
		log.Fatalf("Erro ao dropar tabela: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE produtos (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			nome TEXT NOT NULL,
			preco REAL NOT NULL,
			stock INTEGER NOT NULL
		);
	`)
	if err != nil {
		log.Fatalf("Erro ao criar tabela: %v", err)
	}

	fmt.Println("Populando banco de dados com 50.000 registros...")

	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Erro ao iniciar transação: %v", err)
	}

	stmt, err := tx.Prepare("INSERT INTO produtos (nome, preco, stock) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatalf("Erro ao preparar statement: %v", err)
	}
	defer stmt.Close()

	for i := 1; i <= 50000; i++ {
		nome := fmt.Sprintf("Produto %d", i)
		preco := float64(i) * 1.5
		stock := i % 100

		_, err = stmt.Exec(nome, preco, stock)
		if err != nil {
			tx.Rollback()
			log.Fatalf("Erro ao inserir registro %d: %v", i, err)
		}
	}

	if err = tx.Commit(); err != nil {
		log.Fatalf("Erro ao fazer commit: %v", err)
	}

	fmt.Println("Dataset de 50.000 registros populado com IDs de 1 a 50000!")
}
