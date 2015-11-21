package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"os"
	"strings"
)

var sharedInstance = newDB()

func newDB() *gorm.DB {
	hostname, _ := os.Hostname()
	if strings.Contains(hostname, "local") {
		db, _ := gorm.Open("postgres", "user=rompei dbname=qna sslmode=disable")
		return &db
	}
	db, _ := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	return &db
}

// GetDB returns database object
func GetDB() *gorm.DB {
	return sharedInstance
}
