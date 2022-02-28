package of

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Path   string
	Method string
	Func   gin.HandlerFunc
}

type API interface {
	Register(module Module, handler ...Handler) error
}
