package models

type User struct {
	ID  int64  `db:"id"`
	Sub string `db:"sub"`
}
