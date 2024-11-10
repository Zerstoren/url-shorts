package system

import (
	"github.com/a-h/templ"
	"time"
)

type Response interface {
	SetTitle(title string) Response
	SetKeyword(keyword string) Response
	SetDescription(description string) Response
	SetCacheTime(duration time.Duration) Response

	GetTitle() string
	GetKeyword() string
	GetDescription() string
	GetContent() templ.Component
	GetCacheTime() *time.Duration
	GetRedirect() *string
}

func NewResponseData(content templ.Component) *ResponseData {
	return &ResponseData{
		content:   content,
		code:      200,
		cacheTime: nil,
	}
}

func NewResponseRedirect(target string) *ResponseData {
	return &ResponseData{
		redirect: &target,
		code:     301,
	}
}

type ResponseData struct {
	title       string
	keyword     string
	description string
	content     templ.Component
	cacheTime   *time.Duration
	code        int
	redirect    *string
}

func (r *ResponseData) SetTitle(title string) Response {
	r.title = title
	return r
}

func (r *ResponseData) SetKeyword(keyword string) Response {
	r.keyword = keyword
	return r
}

func (r *ResponseData) SetDescription(description string) Response {
	r.description = description
	return r
}

func (r *ResponseData) SetCacheTime(duration time.Duration) Response {
	r.cacheTime = &duration
	return r
}

func (r *ResponseData) GetTitle() string {
	return r.keyword
}

func (r *ResponseData) GetKeyword() string {
	return r.keyword
}

func (r *ResponseData) GetDescription() string {
	return r.description
}

func (r *ResponseData) GetContent() templ.Component {
	return r.content
}

func (r *ResponseData) GetCacheTime() *time.Duration {
	return r.cacheTime
}

func (r *ResponseData) GetRedirect() *string {
	return r.redirect
}
