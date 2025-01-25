package main

import (
	"fmt"
	"gin-twitter/db"
	"gin-twitter/models"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(
		&models.User{},
		&models.Tweet{},
		&models.Comment{},
		&models.Retweet{},
		&models.Favorite{},
	)
}
