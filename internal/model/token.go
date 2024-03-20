package model

type Token struct {
	ID           int    `db:"id"`
	RefreshToken string `db:"refreshtoken"`
	UserID       int    `db:"user_id"`
}
