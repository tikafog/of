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
	Register(module ModuleStater, handler ...APIHandler) error
}
