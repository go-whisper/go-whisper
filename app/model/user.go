package model

type User struct {
	ID       uint
	Name     string
	Password string
}

func (u User) TableName() string {
	return "users"
}
