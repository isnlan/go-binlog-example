package main

import "time"

type User struct {
	Id      int       `gorm:"column:id"`
	Name    string    `gorm:"column:name"`
	Status  string    `gorm:"column:status"`
	Created time.Time `gorm:"column:created"`
}

func (User) TableName() string {
	return "user"
}

func (User) SchemaName() string {
	return "mydb"
}
