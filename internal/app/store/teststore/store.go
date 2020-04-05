package teststore

import (
	"github.com/Traliaa/http-rest-api/internal/app/model"
	"github.com/Traliaa/http-rest-api/internal/app/store"
)

//Store...
type Store struct {
	userRepository *UserRepository
}

//New store..
func New() *Store {
	return &Store{}
}

//user
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}
	return s.userRepository
}
