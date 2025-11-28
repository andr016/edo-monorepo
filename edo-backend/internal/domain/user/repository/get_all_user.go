package repository

import (
	model "admin-panel/internal/domain/user/repository/models"
	"fmt"
	"github.com/jmoiron/sqlx"

	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/lib/pq"
)

type UserStorage struct {
	db *sqlx.DB
}

func NewUserStorage(db *sqlx.DB) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

func (s *UserStorage) GetAll() ([]model.User, error) {
	query, args, err := dialect.
		Select("login").
		From("users").
		ToSQL()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	var users []model.User
	err = s.db.Select(&users, query, args...)
	return users, err
}
