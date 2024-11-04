package system

import "github.com/a-h/templ"

type Response interface {
	SetTitle(title string) Response
	SetKeyword(keyword string) Response
	SetDescription(description string) Response

	GetTitle() string
	GetKeyword() string
	GetDescription() string
	GetContent() templ.Component
}

func NewResponseData(content templ.Component) *ResponseData {
	return &ResponseData{
		content: content,
	}
}

type ResponseData struct {
	title       string
	keyword     string
	description string
	content     templ.Component
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
