package intchess

import (
	"code.google.com/p/go.crypto/bcrypt"
)

type User struct {
	Id          int64
	Username    string `sql:"type:varchar(60);"`
	AccessToken string `sql:"type:varchar(60);"`
	Email       string `sql:"type:varchar(100);"`
	IsAi        bool
	VersesAi    bool
	CreatedAt   NullTime
	UpdatedAt   NullTime
	DeletedAt   NullTime
}

func AttemptLogin(propUsername string, propPassword string) *User {
	var propUser User
	err := dbGorm.Where(&User{Username: propUsername}).First(&propUser).Error

	if err == nil {
		if err = bcrypt.CompareHashAndPassword([]byte(propUser.AccessToken), []byte(propPassword)); err == nil {
			//they have passed the login check. Save them to the session and redirect to management portal
			return &propUser
		}
	}
	return nil
}
