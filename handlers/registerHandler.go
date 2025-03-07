package handlers

import (
    "net/http"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	"my-go-project/users" 
)


func Register(c *gin.Context) {
		// Открываем базу данных 
		db, err := sql.Open("sqlite3", "users.db")
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()
		

	var newUser users.User

	if err := c.ShouldBindJSON(&newUser); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	if err := users.IsEmpty(newUser.UserName, "Name"); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	if err := users.IsEmpty(newUser.Password, "Password"); err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}

	if err := users.UserExists(newUser.UserName); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
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

	users.UserInsertDB(newUser)

	token, err := users.CreateToken(newUser.UserName, newUser.ID) 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": token})
}
