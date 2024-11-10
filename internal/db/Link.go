package db

type Link struct {
	ID     uint `gorm:"primarykey"`
	Target string
}
