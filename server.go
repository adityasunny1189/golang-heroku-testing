package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type User struct {
	name  string `json:"Name"`
	email string `json:"Email"`
}

var users = []User{}

func AddUser(c *gin.Context) {
	var user User
	user.name = c.PostForm("name")
	user.email = c.PostForm("email")
	log.Println(user)
	users = append(users, user)
	log.Println(users)
}

func main() {
	log.Printf("Backend Testing for corsit")

	r := gin.Default()
	r.POST("/adduser", AddUser)
	r.Run("localhost:5000")
}
