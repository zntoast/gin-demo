package router

import (
	"gindemo/router/system"
)

type RouterGroup struct {
	system.UserRouter
}

var RouterGroupApp = new(RouterGroup)
