package models

import (
	"gorm.io/gorm"
	"strconv"
)

type User struct {
	gorm.Model
	Name       string `json:"name"`
	TelegramId int64  `json:"telegram_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	ChatId     int64  `json:"chat_id"`
}

type UserModel struct {
	Db *gorm.DB
}

func (m *UserModel) Create(user User) error {

	result := m.Db.Create(&user)

	return result.Error
}

func (m *UserModel) FindOne(telegramId int64) (*User, error) {
	existUser := User{}

	result := m.Db.First(&existUser, User{TelegramId: telegramId})

	if result.Error != nil {
		return nil, result.Error
	}

	return &existUser, nil
}

func (m *UserModel) FindAll() ([]User, error) {
	var existUser []User

	result := m.Db.Find(&existUser)

	if result.Error != nil {
		return nil, result.Error
	}

	return existUser, nil
}

func (u User) Recipient() string {
	return strconv.FormatInt(u.ChatId, 10)
}
