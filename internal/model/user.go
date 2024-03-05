package model

type User struct {
	Email    string `db:"email"`
	Password string `db:"password"`
}
