package errors

import (
	"github.com/gin-gonic/gin"
	"github.com/wayt/go-starter/controllers"
	"net/http"
)

type originer interface {
	Origin() error
}

type typer interface {
	Type() error
}

// JSONErr error object
type JSONErr struct {
	Text   string `json:"text"`
	Origin string `json:"origin,omitempty"`
}

func formatJSON(err error) map[string]interface{} {

	origin := ""
	if fct, ok := err.(originer); ok {
		origin = fct.Origin().Error()
	}

	return map[string]interface{}{
		"error": JSONErr{
			Text:   err.Error(),
			Origin: origin,
		},
	}
}

// Handle error reponse
func Handle(c *gin.Context, err error) {

	if fct, ok := err.(typer); ok {

		switch fct.Type() {
		case controllers.ErrDuplicateEntry:
			c.JSON(http.StatusConflict, formatJSON(err))
		case controllers.ErrNotFound:
			c.JSON(http.StatusNotFound, formatJSON(err))
		default:
			c.JSON(http.StatusConflict, formatJSON(err))
		}

	} else {
		c.JSON(http.StatusInternalServerError, formatJSON(err))
	}
}
