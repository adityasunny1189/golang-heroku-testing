package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	name  string `json:"Name"`
	email string `json:"Email"`
}

var users = []User{
	{name: "Aditya Pathak", email: "aditya@gmail.com"},
}

func AddUser(c *gin.Context) {
	var user User
	user.name = c.PostForm("name")
	user.email = c.PostForm("email")
	log.Println(user)
	users = append(users, user)
	log.Println(users)
}

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func main() {
	log.Printf("Backend Testing for corsit")

	r := gin.Default()
	r.GET("/", GetUsers)
	r.POST("/adduser", AddUser)
	r.Run("localhost:5000")
}
