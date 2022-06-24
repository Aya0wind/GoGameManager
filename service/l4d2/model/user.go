package model

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int64  `gorm:"column:id,primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
}

func (User) TableName() string {
	return "user"
}
