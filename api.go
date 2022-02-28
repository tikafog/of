package of

import (
	"github.com/gin-gonic/gin"
)

type API interface {
	Register(path string, method string, fn gin.HandlerFunc) error
}
