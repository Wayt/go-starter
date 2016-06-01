package ponies

import (
	"github.com/gin-gonic/gin"
	"github.com/wayt/go-starter/api/v1/errors"
	"github.com/wayt/go-starter/controllers"
	"net/http"
)

// GetPonies return all ponies
func GetPonies(c *gin.Context) {

	ponies, err := controllers.Ponies.List()
	if err != nil {
		errors.Handle(c, err)
		return
	}

	c.JSON(http.StatusOK, ponies)
}

// GetPony return one pony, by id
func GetPony(c *gin.Context) {

	id := c.Param("id")

	pony, err := controllers.Ponies.Get(id)
	if err != nil {
		errors.Handle(c, err)
		return
	}

	c.JSON(http.StatusOK, pony)
}

// PostPony create one pony
func PostPony(c *gin.Context) {

	form := struct {
		Name  string `form:"name" json:"name" binding:"required"`
		Title string `form:"title" json:"title"`
	}{}

	if err := c.Bind(&form); err != nil {
		errors.Handle(c, err)
		return
	}

	pony, err := controllers.Ponies.Create(form.Name, form.Title)
	if err != nil {
		errors.Handle(c, err)
		return
	}

	c.JSON(http.StatusCreated, pony)
}

// DeletePony delete one pony
func DeletePony(c *gin.Context) {

	id := c.Param("id")

	pony, err := controllers.Ponies.Get(id)
	if err != nil {
		errors.Handle(c, err)
		return
	}

	pony, err = controllers.Ponies.Delete(pony)
	if err != nil {
		errors.Handle(c, err)
		return
	}

	c.JSON(http.StatusAccepted, pony)
}
