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
		log.Fatal(err)
	}
	defer db.Close()

	var id int
	err = db.QueryRow("SELECT id FROM produtos ORDER BY id ASC LIMIT 1").Scan(&id)
	if err != nil {
		log.Fatalf("Erro ao buscar ID: %v", err)
	}
	fmt.Printf("O menor ID existente no banco é: %d\n", id)
}
