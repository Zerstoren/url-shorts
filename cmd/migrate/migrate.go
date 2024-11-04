package main

import (
	"fmt"
	"url-shorts.com/internal/db"
)

func main() {
	dba := db.GetDb()

	fmt.Printf("Migrate `%s` completed with: %v\n", "link", dba.AutoMigrate(db.Link{}))
	fmt.Println("Migration completed")
}
