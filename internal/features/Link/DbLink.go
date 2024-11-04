package featureLink

import (
	"url-shorts.com/internal/db"
)

type linkMethods interface {
	db.Iterable[LinkItem]
	createLink(url string) error
	getByCode(code string) error
}

type Link interface {
	GetId() uint
	GetTarget() string
	GetShortUrl() string
}

func newLinkRequest() linkMethods {
	return &linkRequest{
		IterableOrigin: db.IterableOrigin[LinkItem]{
			Origin: &[]LinkItem{},
		},
		Request: db.Request{
			Db: db.GetDb(),
		},
	}
}

type linkRequest struct {
	db.IterableOrigin[LinkItem]
	db.Request
}

func (l *linkRequest) getByCode(code string) error {
	id := idFromCode(code)
	return l.Db.First(l.Origin, id).Error
}

func (l *linkRequest) createLink(url string) error {
	link := &[]LinkItem{{
		db.Link{
			Target: url,
		},
	}}

	l.Origin = link

	return l.Db.Save(&l.Origin).Error
}

type LinkItem struct {
	db.Link
}

func (l *LinkItem) TableName() string {
	return "links"
}

func (l *LinkItem) GetId() uint {
	return l.ID
}

func (l *LinkItem) GetTarget() string {
	return l.Target
}

func (l *LinkItem) GetShortUrl() string {
	return codeFromId(l.ID)
}
