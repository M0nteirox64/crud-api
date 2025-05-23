package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string `json:"Name"`
	DisplayName string `json:"DisplayName"`
}

var currentUser *User

func main() {
	r := gin.Default()

	r.GET("/", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Message": "API Online!",
		})
	})

	r.GET("/user", func (c *gin.Context) {
		if currentUser == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"Status": "No accounts created.",
			})
			return
		}
		
		c.JSON(http.StatusOK, currentUser)
	})

	r.POST("/user", func (c *gin.Context) {
		var newUser User

		if err := c.ShouldBindJSON(&newUser); err != nil { // Vai pegar o conteudo do request post e mandar pra instancia "newUser" [struct]
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Invalid JSON.",
			})
			return
		}

		currentUser = &newUser
		c.JSON(http.StatusCreated, gin.H{
			"status": "Created",
			"profile": currentUser,
		})
	})
	r.Run()
}
