package resolver

import (
	"errors"
	"fmt"
)

package main

import (
"errors"
"fmt"

"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID
	Name  string
	Email string
}

type Resolver struct {
	users []User
}

func (r *Resolver) User(args struct{ ID uuid.UUID }) (*User, error) {
	for _, u := range r.users {
		if u.ID == args.ID {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *Resolver) Users() []User {
	return r.users
}

func (r *Resolver) CreateUser(args struct{ Input UserInput }) (*User, error) {
	u := User{
		ID:    uuid.New(),
		Name:  args.Input.Name,
		Email: args.Input.Email,
	}
	r.users = append(r.users, u)
	return &u, nil
}

func (r *Resolver) UpdateUser(args struct{ ID uuid.UUID, Input UserInput }) (*User, error) {
	for i, u := range r.users {
		if u.ID == args.ID {
			r.users[i].Name = args.Input.Name
			r.users[i].Email = args.Input.Email
			return &r.users[i], nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *Resolver) DeleteUser(args struct{ ID uuid.UUID }) (bool, error) {
	for i, u := range r.users {
		if u.ID == args.ID {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("user not found")
}

func NewResolver() *Resolver {
	return &Resolver{
		users: []User{
			{ID: uuid.New(), Name: "Alice", Email: "alice@example.com"},
			{ID: uuid.New(), Name: "Bob", Email: "bob@example.com"},
		},
	}
}

func (u User) String() string {
	return fmt.Sprintf("{ID: %s, Name: %s, Email: %s}", u.ID.String(), u.Name, u.Email)
}

func (ui UserInput) String() string {
	return fmt.Sprintf("{Name: %s, Email: %s}", ui.Name, ui.Email)
}