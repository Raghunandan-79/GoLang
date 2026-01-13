package user

import "fmt"

type user struct {
	name  string
	email string
}

func New(name string, email string) *user {
	return &user{
		name: name,
		email: email,
	}
}

func (u *user) SetDetails(name string, email string) {
	u.name = name
	u.email = email
}

func (u *user) GetDetails() {
	fmt.Println("Name:", u.name, "\nEmail:", u.email)
}