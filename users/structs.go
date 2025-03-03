package users 

import (
    _ "github.com/mattn/go-sqlite3"
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID  	 int	`jsom:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}