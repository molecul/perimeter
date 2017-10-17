package web

import "github.com/gin-gonic/gin"

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

type Link struct {
	Path string
	Name string
	Icon string
}

type Routes []Route

type Links []Link
