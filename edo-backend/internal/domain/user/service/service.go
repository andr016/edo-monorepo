package service

import "admin-panel/internal/domain/user/repository"

type Service struct {
	repository repository.Repository
}
