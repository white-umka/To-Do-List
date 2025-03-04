package handlers 

import (
    "net/http"
	"database/sql"
	"fmt"
	"io"

	"my-go-project/users" 
    
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)


func Login(c *gin.Context) {
	// Открываем базу данных 
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	
	
	var newUser users.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Логируем тело запроса
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Request body:", string(body))

	
	if err := users.IsEmpty(newUser.UserName, "Name"); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	if err := users.IsEmpty(newUser.Password, "Password"); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	if err := users.LenName(newUser.UserName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := users.LenPassword(newUser.Password); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	if err := users.UserLogin(db, newUser.UserName, newUser.Password); err!= nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully!"})
}