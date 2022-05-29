package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

var users = []User{
	{Name: "Aditya Pathak", Email: "aditya@gmail.com"},
}

func AddUser(c *gin.Context) {
	var user User
	user.Name = c.PostForm("name")
	user.Email = c.PostForm("email")
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
	r.Run(":5000")
}
