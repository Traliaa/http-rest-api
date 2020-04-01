package store_test

import (
	"github.com/Traliaa/http-rest-api/internal/app/model"
	"github.com/Traliaa/http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "test@test.ru",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
