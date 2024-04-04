package model

type Token struct {
	ID           int    `db:"id"`
	RefreshToken string `db:"refresh_token"`
	UserID       int    `db:"user_id"`
}
