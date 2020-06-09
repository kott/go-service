package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/kott/go-service/errors"
)

const (
	jsonHeader = "application/json"
)

var allowedContentTypes = map[string]interface{}{
	"application/json":               nil,
	"application/json;charset=utf-8": nil,
}

// JSONResponseHeader sets the content-type of responses to application/json.
func JSONResponseHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", jsonHeader)
		c.Next()
	}
}

// ForceJSON will require the request to have a json content-type
func ForceJSON() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := strings.TrimSpace(strings.ToLower(c.ContentType()))
		if _, ok := allowedContentTypes[header]; !ok && containsBody(c.Request.Method) {
			c.AbortWithStatusJSON(http.StatusUnsupportedMediaType, errors.NewAppError(errors.UnsupportedMediaType,
				errors.Descriptions[errors.UnsupportedMediaType], "nil"))
		}
		c.Next()
	}
}

func containsBody(method string) bool {
	switch method {
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		return true
	default:
		return false
	}
}
