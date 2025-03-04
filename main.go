package main

import (
    "database/sql"

    "my-go-project/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
    // <-------------------------PREPARATION---------------------------->
    gin.SetMode(gin.DebugMode)
    
    db, err := sql.Open("sqlite3", "users.db")
    if err != nil {
      panic(err.Error())
    }
    defer db.Close()

	r := gin.Default()

    //<-------------------------STATIC-FILES---------------------------->
    r.Static("/Front-end", "./Front-end")
    r.LoadHTMLGlob("Front-end/templates/*")

    //<-------------------------START-SERVER---------------------------->
    
    r.POST("/register", handlers.Register) // http://localhost:8080/register
    r.GET("/register", func(c *gin.Context){
        c.HTML(200, "register.html", nil)
    })

    r.POST("/login", gin.Logger(), handlers.Login) // http://localhost:8080/login
    r.GET("/login", func(c *gin.Context){
        c.HTML(200, "login.html", nil)
    })

    r.Run()
}


