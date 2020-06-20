package deepl

import "strings"

type Requester interface {
	SetAuth(key string)
	Path() string
	Method() string
	Query() string
	Param() *strings.Reader
}
