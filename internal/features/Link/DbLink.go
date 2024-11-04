package featureLink

import (
	"gorm.io/gorm"
	"url-shorts.com/internal/system"
)

type linkMethods interface {
	getFirst() (Link, bool)
	createLink(url string, short string) error
	findLastItem() error
}

type Link interface {
	GetId() uint
	GetTarget() string
	GetShortUrl() string
}

func newLinkRequest() linkMethods {
	return &linkRequest{
		Origin: nil,
		DbRequest: system.DbRequest{
			Db: system.GetDb(),
		},
	}
}

type linkRequest struct {
	Origin *[]LinkItem
	system.DbRequest
}

func (l *linkRequest) findLastItem() error {
	return l.Db.Order("id DESC").Limit(1).Find(&l.Origin).Error
}

func (l *linkRequest) getFirst() (Link, bool) {
	if l.Origin == nil {
		return nil, false
	}

	if len(*l.Origin) == 0 {
		return nil, false
	}

	return &(*l.Origin)[0], true
}

func (l *linkRequest) createLink(url string, short string) error {
	link := &[]LinkItem{{
		Target: url,
		Short:  short,
	}}

	l.Origin = link

	return l.Db.Save(l.Origin).Error
}

type LinkItem struct {
	*gorm.Model
	Target string
	Short  string `gorm:"uniqueIndex"`
}

func (l *LinkItem) GetId() uint {
	return l.ID
}

func (l *LinkItem) GetTarget() string {
	return l.Target
}

func (l *LinkItem) GetShortUrl() string {
	return l.Short
}
