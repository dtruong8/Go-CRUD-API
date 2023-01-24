package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
}

type UserService interface {
	GetUser(id int) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id int) error
}
type userService struct {
	users map[int]*User
}

func (s *userService) GetUser(id int) (*User, error) {
	user, ok := s.users[id]
	if !ok {
		return nil, fmt.Errorf("user with id %d not found", id)
	}
	return user, nil
}

func (s *userService) CreateUser(user *User) error {
	if _, ok := s.users[user.ID]; ok {
		return fmt.Errorf("user with id %d already exists", user.ID)
	}
	s.users[user.ID] = user
	return nil
}

func (s *userService) UpdateUser(user *User) error {
	if _, ok := s.users[user.ID]; !ok {
		return fmt.Errorf("user with id %d not found", user.ID)
	}
	s.users[user.ID] = user
	return nil
}

func (s *userService) DeleteUser(id int) error {
	if _, ok := s.users[id]; !ok {
		return fmt.Errorf("user with id %d not found", id)
	}
	delete(s.users, id)
	return nil
}

func main() {
	// Initialize userService
	service := &userService{
		users: make(map[int]*User),
	}
	// Example usage
	user1 := &User{ID: 1, Name: "John Doe", Email: "johndoe@example.com"}
	err := service.CreateUser(user1)
	if err != nil {
		fmt.Println(err)
	}
	user, err := service.GetUser(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Name:", user.Name)
	}
	user1.Name = "Jane Doe"
	err = service.UpdateUser(user1)
	if err != nil {
		fmt.Println(err)
	}
	user, err = service.GetUser(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Name:", user.Name)
	}
	err = service.DeleteUser(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Deleted #id 1")
	}
}
