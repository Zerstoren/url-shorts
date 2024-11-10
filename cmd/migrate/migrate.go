package main

import (
	"fmt"
	"url-shorts.com/internal/db"
)

func main() {
	dba := db.GetDb()

	fmt.Printf("Migrate `%s` completed with: %v\n", "link", dba.AutoMigrate(db.Link{}))
	fmt.Printf("Migrate `%s` completed with: %v\n", "user", dba.AutoMigrate(db.User{}))
	fmt.Println("Migration completed")
}
