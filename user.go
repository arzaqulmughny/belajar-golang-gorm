package belajargolanggorm

import "time"

type User struct {
	Id int `gorm:"primary_key;column:id;<-:create"`
	Name string `gorm:"column:name"`
	Email string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	updatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (u *User) TableName() string {
	return "users"
}