package users 

import (
    _ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID  	 int	`jsom:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

