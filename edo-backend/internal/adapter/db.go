package repository

import (
	"admin-panel/internal/config"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

type Provider interface {
	ConnectDB() *sqlx.DB
}

type provider struct{}

func NewProvider() Provider {
	return &provider{}
}

func (p *provider) ConnectDB() *sqlx.DB {
	dsn, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Open("pgx", dsn.TmpPath)
	if err != nil {
		log.Fatalf("Ошибка при открытии соединения: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Не удалось подключиться к БД: %v", err)
	}

	log.Println("Подключение к PostgreSQL успешно через sqlx!")
	return db
}
