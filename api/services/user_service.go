package services

import (
	"errors"
	"sync"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserService struct {
	users map[string]User
	mutex sync.RWMutex
}

func NewUserService() *UserService {
	return &UserService{
		users: make(map[string]User),
	}
}

func (s *UserService) GetUsers() ([]User, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	users := make([]User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users, nil
}

func (s *UserService) GetUser(id string) (*User, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	user, exists := s.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (s *UserService) CreateUser(user *User) (*User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// In a real application, you would generate a proper ID
	// and validate the user data
	if user.ID == "" {
		return nil, errors.New("user ID is required")
	}

	if _, exists := s.users[user.ID]; exists {
		return nil, errors.New("user already exists")
	}

	s.users[user.ID] = *user
	return user, nil
}
