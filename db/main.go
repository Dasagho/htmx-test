package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	logging "github.com/dasagho/htmx-test/log"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

var credentials = map[string]string{
	"host":   "",
	"port":   "",
	"user":   "",
	"pass":   "",
	"dbname": "",
	"ssl":    "",
}

var defaultCredentials = map[string]string{
	"host":   "localhost",
	"port":   "5432",
	"user":   "postgres",
	"pass":   "root",
	"dbname": "database_test",
	"ssl":    "disable",
}

func ConnectDB() {
	initializeCredentials()
	// Connect to check table exists
	db, err = sql.Open("postgres", fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=%s`,
		credentials["host"], credentials["port"], credentials["user"],
		credentials["pass"], "postgres", credentials["ssl"]))
	if err != nil {
		logging.Error("Failed to connect database")
		log.Println(err)
		return
	}
	logging.Debug("Conectado a base de datos 0...")
	logging.Debug("Checkeando existencia de base de datos...")

	// Ejecutar script para crear la base de datos y cerrar conexion
	err = executeSQLFile(db, filepath.Join("db", "migrations", "0-createDataBase.sql"))
	if err != nil {
		logging.Error("Failed to execute migration 0")
		log.Println(err)
	}
	db.Close()

	// Connect to correctDatabase
	db, err = sql.Open("postgres", fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=%s`,
		credentials["host"], credentials["port"], credentials["user"],
		credentials["pass"], credentials["dbname"], credentials["ssl"]))
	if err != nil {
		logging.Error("Failed to connect database")
		log.Println(err)
		return
	}

	logging.Debug("Ejecutando migraciones...")
	// Ejecutar scripts para borrar y crear tablas
	err = executeSQLFile(db, filepath.Join("db", "migrations", "1-deleteTables.sql"))
	if err != nil {
		logging.Error("Failed to execute migration 1")
		log.Println(err)
		return
	}

	err = executeSQLFile(db, filepath.Join("db", "migrations", "2-createTables.sql"))
	if err != nil {
		logging.Error("Failed to execute migration 2")
		log.Println(err)
		return
	}
}

func initializeCredentials() {
	for key := range credentials {
		value, err := os.LookupEnv(key)
		if err {
			credentials[key] = value
		} else {
			credentials[key] = defaultCredentials[key]
		}
	}
}

func executeSQLFile(db *sql.DB, filePath string) error {
	content, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		log.Printf("Error leyendo el archivo SQL: %v", err)
		return err
	}

	_, err = db.Exec(string(content))
	if err != nil {
		log.Printf("Error ejecutando el archivo SQL (%s): %v", filePath, err)
		return err
	}
	return nil
}
