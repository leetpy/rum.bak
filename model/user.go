package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Id       uint64
	Email    string
	Password string
}

func (u *User) GetUserById(id uint64) *User {
	user := User{}
	return &user
}

func (u *User) AddUser(user *User) {
	db.Create(user)
}

func (u *User) RemoveUser(user *User) {
	db.Delete(user)
}
