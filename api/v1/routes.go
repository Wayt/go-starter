package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/wayt/go-starter/api/v1/ponies"
)

// Init initialize v1 router group
func Init(r *gin.RouterGroup) {

	r.GET("/ponies", ponies.GetPonies)
	r.GET("/ponies/:id", ponies.GetPony)
	r.POST("/ponies", ponies.PostPony)
	r.DELETE("/ponies/:id", ponies.DeletePony)
}
