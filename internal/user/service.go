package user

import (
	"errors"
	"strings"
)

var (
	ErrNotFound = errors.New("user not found")
	ErrBadInput = errors.New("invalid input")
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ListUsers() []User {
	return s.repo.List()
}

func (s *Service) GetUserByID(id string) (User, error) {
	u, ok := s.repo.GetByID(id)
	if !ok {
		return User{}, ErrNotFound
	}
	return u, nil
}

func (s *Service) CreateUser(name string) (User, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return User{}, ErrBadInput
	}
	return s.repo.Create(name), nil
}

func (s *Service) DeleteUser(id string) error {
	if ok := s.repo.Delete(id); !ok {
		return ErrNotFound
	}
	return nil
}