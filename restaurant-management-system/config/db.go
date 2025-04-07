package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq" // Importa el driver de PostgreSQL para database/sql
)

// ConnectGORM establece la conexión a la base de datos PostgreSQL usando GORM
func ConnectGORM() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos con GORM: %v", err)
		return nil, err
	}

	log.Println("Conexión a la base de datos PostgreSQL establecida correctamente con GORM")
	return db, nil
}

// ConnectSQL establece la conexión a la base de datos PostgreSQL usando database/sql
func ConnectSQL(dbName string) (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := dbName

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos con database/sql: %v", err)
		return nil, err
	}

	// Verifica la conexión
	if err := db.Ping(); err != nil {
		log.Fatalf("No se pudo establecer conexión con la base de datos usando database/sql: %v", err)
		return nil, err
	}

	log.Println("Conexión a la base de datos PostgreSQL establecida correctamente con database/sql")
	return db, nil
}
