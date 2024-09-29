package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LatName   string `json:"lat_name"`
	Email     string `json:"email"`
	Password  []byte `json:"-"`
	Phone     string `json:"phone"`
}

func (user *User) SetPassword(password string) {
	bashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = bashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
