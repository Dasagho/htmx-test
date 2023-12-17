package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	logging "github.com/dasagho/htmx-test/log"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var (
	db          *sql.DB
	err         error
	credentials = map[string]string{
		"DB_HOST": "",
		"DB_PORT": "",
		"DB_USER": "",
		"DB_PASS": "",
		"DB_NAME": "",
		"DB_SSL":  "",
	}

	defaultCredentials = map[string]string{
		"DB_HOST": "localhost",
		"DB_PORT": "5432",
		"DB_USER": "postgres",
		"DB_PASS": "root",
		"DB_NAME": "database_test",
		"DB_SSL":  "disable",
	}

	errDBAlreadyExists = errors.New("error database already exists")
)

func ConnectDB() {
	initializeCredentials()

	// Connect to Database
	connectionCredentials := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=%s`,
		credentials["DB_HOST"], credentials["DB_PORT"], credentials["DB_USER"],
		credentials["DB_PASS"], "postgres", credentials["DB_SSL"])

	err = connectWithRetry(connectionCredentials, 5)

	if err != nil {
		logging.Error("Fallo al intentar conectar a base de datos con las siguientes credenciales" + connectionCredentials)
		log.Printf("Fallo al intentar conectar a base de datos con las siguientes credenciales: %s", connectionCredentials)
		return
	}

	// check if table exists
	logging.Debug("Conectado a base de datos 0...")
	logging.Debug("Checkeando existencia de base de datos...")

	// Ejecutar script para crear la base de datos y cerrar conexion
	err = executeSQLFile(db, filepath.Join("db", "migrations", "0-createDataBase.sql"))
	if err != nil && !errors.Is(err, errDBAlreadyExists) {
		logging.Error("Failed to execute migration 0")
		log.Println(err)
	}
	db.Close()

	// Connect to correctDatabase
	db, err = sql.Open("postgres", fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=%s`,
		credentials["DB_HOST"], credentials["DB_PORT"], credentials["DB_USER"],
		credentials["DB_PASS"], credentials["DB_NAME"], credentials["DB_SSL"]))
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

func connectWithRetry(dbURL string, maxAttempts int) error {
	var err error

	for attempts := 0; attempts < maxAttempts; attempts++ {
		db, err = sql.Open("postgres", dbURL)
		if err != nil {
			logging.Error("No se pudo conectar a la base de datos: " + err.Error() + ". Reintentando...")
			log.Printf("No se pudo conectar a la base de datos: %v. Reintentando...", err)
		} else {
			err = db.Ping()
			if err == nil {
				logging.Debug("Conexión a la base de datos establecida con éxito.")
				log.Println("Conexión a la base de datos establecida con éxito.")
				return nil
			}
			log.Printf("Error al realizar ping a la base de datos: %v. Reintentando...", err)
		}

		// Esperar antes del próximo reintento
		time.Sleep(2 * time.Second)
	}

	return fmt.Errorf("después de %d intentos, no se pudo conectar a la base de datos: %v", maxAttempts, err)
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
		if pqErr, ok := err.(*pq.Error); ok {
			// Verificar el código de error específico de PostgreSQL
			if pqErr.Code == "42P04" {
				// Manejar el error específico aquí
				logging.Debug("Database already Exists")
				return errDBAlreadyExists
			}
		} else {
			log.Printf("Error ejecutando el archivo SQL (%s): %v", filePath, err)
			return err
		}
	}
	return nil
}
