package main

import (
	"fmt"
	featureLink "url-shorts.com/internal/features/Link"
	"url-shorts.com/internal/system"
)

func main() {
	db := system.GetDb()

	fmt.Printf("Migrate `%s` completed with: %v\n", "link", db.AutoMigrate(featureLink.LinkItem{}))
	fmt.Println("Migration completed")
}
