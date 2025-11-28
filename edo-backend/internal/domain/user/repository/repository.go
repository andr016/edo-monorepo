package repository

import (
	db "admin-panel/internal/adapter"
	"github.com/doug-martin/goqu/v9"
)

var dialect = goqu.Dialect("postgres")

type Repository struct {
	dbProvider db.Provider
}

func New(dbProvider db.Provider) *Repository {
	return &Repository{dbProvider: dbProvider}
}
