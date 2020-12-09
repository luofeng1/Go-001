package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luofeng1/Go-000/Week02/homework/code"
	"github.com/luofeng1/Go-000/Week02/homework/service"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/user/:id", func(c *gin.Context) {
		userID, _ := c.Params.Get("id")
		user, err := service.GetUser(userID)

		if err != nil {
			var e *code.StatusError
			if errors.Is(err, e) {
				c.JSON(http.StatusOK, e)
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"err": "unKnow",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": user,
		})
	})

	_ = r.Run(":10086")
}
