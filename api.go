package of

import (
	"github.com/gin-gonic/gin"
)

type APIHandler struct {
	Path   string
	Method string
	Func   gin.HandlerFunc
}

type API interface {
	SetRoots(API)
	Register(module ModuleStarter, handler ...APIHandler) error
}
