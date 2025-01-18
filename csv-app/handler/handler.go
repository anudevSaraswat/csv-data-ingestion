package handler

import (
	"csv-app/models"
	"csv-app/schema"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func APIGetUser(c *gin.Context) {

	db := c.MustGet("db").(*sql.DB)

	var payload models.User
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, &models.Response{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: err.Error(),
			Data:         []interface{}{},
		})
		return
	}

	// TODO:
	// check for the record in cache
	// if cache hit, then return the data
	// if cache miss, fetch it from the database and store it in the cache

	users, err := schema.FetchUsers(db, payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &models.Response{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "Something went wrong",
			Data:         []interface{}{},
		})
		return
	}

	c.JSON(http.StatusInternalServerError, &models.Response{
		StatusCode:   http.StatusOK,
		ErrorMessage: "",
		Data:         users,
	})

}
