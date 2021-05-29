package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type handlerFunc func(c *gin.Context) (obj interface{}, err error)

func JSON(f handlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		obj, err := f(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if obj == nil {
			c.Status(http.StatusOK)
			return
		}

		c.JSON(http.StatusOK, obj)
	}
}
