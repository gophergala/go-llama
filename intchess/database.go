package intchess

import (
	"code.google.com/p/go.crypto/bcrypt"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var dbGorm *gorm.DB

func ConnectToDatabase() error {
	dbGormConnection, err := gorm.Open("mysql", "intchess:lolCAKEolol123@tcp(106.187.103.125:3306)/intchess")
	if err != nil {
		return err
	}
	dbGorm = &dbGormConnection
	return nil
}

func CreateDatabaseTables() {
	fmt.Printf("I am attempting to create the database tables!")
	fmt.Printf("Dropping (if exists) and creating Users table...\n")
	dbGorm.DropTableIfExists(&User{})
	dbGorm.CreateTable(&User{})
	//create me a default user
	fmt.Printf("Adding default test user to database...\n")
	pass, _ := bcrypt.GenerateFromPassword([]byte("test"), 3)
	u := User{
		Username:    "test",
		AccessToken: string(pass),
		IsAi:        false,
		VersesAi:    true,
	}
	dbGorm.Create(&u)
	fmt.Printf("Done!")
}
