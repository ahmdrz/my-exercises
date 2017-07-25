package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Name      string
	Family    string
	BirthDate time.Time
}

func main() {
	db, err := gorm.Open("sqlite3", "database.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}

	user := User{}
	user.Name = "ahmadreza"
	user.Family = "zibaei"
	user.BirthDate = time.Now()
	db.Model(&user).Create(&user)

	user.BirthDate = time.Now()
	db.Model(&user).Save(&user)
}
