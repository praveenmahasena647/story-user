package postgres

import (
	"database/sql"
)

type User struct {
	ID       uint   `db:id`
	Name     string `db:name`
	EmailID  string `db:emailid`
	Password []byte `db:password`
}

func NewUser(n, e string, p []byte) *User {
	return &User{
		Name:     n,
		EmailID:  e,
		Password: p,
	}
}

func (u *User) Insert() (sql.Result, error) {
	return db.Exec(`INSERT INTO users(name, emailid, password) VALUES ($1, $2 , $3);`, u.Name, u.EmailID, u.Password)
}

func FetchUser(email string) (User, error) {
	var u = User{}
	var e = db.QueryRow(`SELECT emailid, password, name, id FROM users WHERE emailid=$1`, email).Scan(&u.EmailID, &u.Password, &u.Name, &u.ID)
	return u, e
}

func U() *User {
	return &User{}
}

func (u *User) FetchByEmail(emailID string) error {
	return db.QueryRow(`SELECT emailid, id FROM users WHERE emailid=$1`, emailID).Scan(&u.EmailID, &u.ID)
}
