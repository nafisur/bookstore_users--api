package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nafisur/bookstore_users-api/domain/users"
	"github.com/nafisur/bookstore_users-api/services"
	"io/ioutil"
	"net/http"
)

func GetUser(c *gin.Context) {

	c.String(http.StatusNotImplemented, "implement me!")
}

func CreateUser(c *gin.Context) {
	var user users.User
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//Handle error
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		//handle json error
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//handle user creation error
		return
	}

	c.JSON(http.StatusCreated, result)

}

func FindUser(c *gin.Context) {

	c.String(http.StatusNotImplemented, "implement me!")
}

