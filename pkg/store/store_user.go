package store

import (
	"fmt"

	"github.com/NuitdelinfoLesCools/back-end/pkg/store/model"
	"github.com/medtune/go-utils/crypto"
)

type userStore interface {
	CreateUser(string, string, string) error
	AuthentificateUser(string, string) (bool, *model.User, error)
}

// CreateUser .
func (s *Store) CreateUser(email, username, password string) error {
	user := model.User{
		Email:    email,
		Username: username,
		Password: crypto.Sha256(password),
	}

	// validate user
	v, err := s.Valid(user)
	if err != nil || !v {
		return err
	}

	// insert user
	if _, err := s.Insert(&user); err != nil {
		return err
	}

	return nil
}

// GetUser select a single user by it username
func (s *Store) GetUser(username string) (*model.User, error) {
	user := &model.User{}
	has, err := s.Where("username = ?", username).Get(user)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, fmt.Errorf("record doesnt exist")
	}
	return user, nil
}

// GetUserByEmail select user by it email
func (s *Store) getUserByEmail(email string) (*model.User, error) {
	user := &model.User{}
	has, err := s.Where("email = ?", email).Get(user)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, fmt.Errorf("record doesnt exist")
	}
	return user, nil
}

// AuthentificateUser auth a user
func (s *Store) AuthentificateUser(username, password string) (bool, *model.User, error) {
	user, err := s.GetUser(username)
	if err != nil {
		// Database server error or record not found
		return false, nil, fmt.Errorf("username or password incorrect")
	}

	// Hashpassword
	if crypto.Sha256(password) == user.Password {
		return true, user, nil
	}

	// Password is incorrect
	return false, nil, fmt.Errorf("username or password incorrect")
}
