package models

import "fmt"

type User struct {
	Username string `gorm:"size:255; not null;" json:"username"`
	Password string `gorm:"size:255; not null;" json:"password"`
}

func (u *User) CreateUser() (*User, error) {
	fmt.Println("creating new user")
	fmt.Println(DB)
	if err := DB.Create(&u).Error; err != nil {
		return &User{}, err
	}

	return u, nil
}
