package db

import (
	"database/sql"
	"io"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbName   = "db/tasks.db"
	dbScript = "db/task.sql"
)

// OpenDB abre o banco de dados SQLite
func OpenDB() (*sql.DB, error) {
	// Abre uma conexão com o banco de dados SQLite
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}

	// Abre o arquivo SQL
	sqlFile, err := os.Open(dbScript)
	if err != nil {
		return nil, err
	}
	defer sqlFile.Close()

	// Lê o conteúdo do arquivo SQL
	var sb strings.Builder
	if _, err := io.Copy(&sb, sqlFile); err != nil {
		return nil, err
	}

	// Separa as instruções SQL do arquivo
	sqlStatements := strings.Split(sb.String(), ";")

	// Executa as instruções SQL do arquivo no banco de dados
	for _, sqlStatement := range sqlStatements {
		if _, err := db.Exec(sqlStatement); err != nil {
			return nil, err
		}
	}

	return db, nil
}

// CloseDB fecha o banco de dados SQLite
func CloseDB(db *sql.DB) {
	db.Close()
}
