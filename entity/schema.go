package entity

import (
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Name string `valid:"required~Name cannot blank"`
	Url  string `gorm:"uniqueIndex" valid:"url~Url Invalid"`
}

type User struct {
	gorm.Model
	Name        string `valid:"required~Name cannot blank"`
	Email       string `valid:"email~Email invalid format"`
	StudentID   string `valid:"matches(^[B]\\d{7}$)~StudentID invalid format"`
	PhoneNumber string `valid:"matches(^([+]66|0)[0-9]{9}$)~PhoneNumber invalid format"`
	Password    string `valid:"minstringlength(8)~Password must be between 8 and 20 characters, maxstringlength(20)~Password must be between 8 and 20 characters"`
}
