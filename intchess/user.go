package intchess

import (
	"code.google.com/p/go.crypto/bcrypt"
	"fmt"
)

type User struct {
	Id          int64
	Username    string   `sql:"type:varchar(60);unique" json:"username"`
	AccessToken string   `sql:"type:varchar(60);" json:"-"`
	CurrentRank int      `json:"current_rank"`
	IsAi        bool     `json:"is_ai"`
	VersesAi    bool     `json:"verses_ai"`
	CreatedAt   NullTime `sql:"type:datetime" json:"created_at"`
	UpdatedAt   NullTime `sql:"type:datetime" json:"updated_at"`
	DeletedAt   NullTime `sql:"type:datetime" json:"-"`
}

func AttemptLogin(propUsername string, propPassword string) *User {
	var propUser User
	err := dbGorm.Where(&User{Username: propUsername}).First(&propUser).Error

	if err == nil {
		if err = bcrypt.CompareHashAndPassword([]byte(propUser.AccessToken), []byte(propPassword)); err == nil {
			//they have passed the login check.
			return &propUser
		}
	}
	return nil
}

func AttemptCreate(propUsername string, propPassword string) *User {
	var propUser User

	if dbGorm.Where(&User{Username: propUsername}).First(&propUser).RecordNotFound() {
		propUser.Username = propUsername
		hashpass, _ := bcrypt.GenerateFromPassword([]byte(propPassword), 3)
		//if err != nil {
		propUser.AccessToken = string(hashpass)
		dbGorm.Create(&propUser)
		return &propUser
		// } else {
		// 	fmt.Println("Error with bcrypt: " + err.Error())
		// 	return nil
		// }
	} else {
		fmt.Println("Username was taken.")
	}
	return nil
}
