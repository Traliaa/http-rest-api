package store

import "github.com/Traliaa/http-rest-api/internal/app/model"

type UserRepository struct {
	store *Store
}

//Create..
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users(email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email, u.Encrypted_password,
	).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

//find User
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}
