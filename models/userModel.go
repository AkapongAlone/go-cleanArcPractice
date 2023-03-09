package models

import "time"

type User struct {
	ID        int `form:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserName  string		`form:"username"`
	Password  string		`form:"password"`
	CreatedAt time.Time `form:"created_at"`
	UpdatedAt time.Time `form:"updated_at"`
}