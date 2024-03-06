package model

type User struct {
	ID       uint32 `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
