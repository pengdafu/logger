package logger

import (
	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

type BTHandle func(c *Context)
