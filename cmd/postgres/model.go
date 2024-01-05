package postgres

import "database/sql"

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(name, email, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}

func (u *User) Insert() (*sql.Rows, error) {
	return db.Query(`INSERT INTO users (name,email,password) VALUES ($1,$2,$3);`, u.Name, u.Email, u.Password)
}
