package config

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	// Força o uso do arquivo trabalho.db na raiz do projeto
	dbPath, err := filepath.Abs("./trabalho.db")
	if err != nil {
		return nil, err
	}

	log.Printf("Conectando ao banco SQLite em: %s", dbPath)

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir banco sqlite: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("erro ao testar conexão sqlite: %w", err)
	}

	return db, nil
}
