package service

import (
	"github.com/go-martini/martini"
)

func NewServer(port string) {
	m := martini.Classic()
	m.Get("/hello/:name", func(params martini.Params) string {
		return "Hello"
	})
	m.RunOnAddr(":"+port)	
}
