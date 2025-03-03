package users

import (
	"database/sql"
    "fmt"

    _ "github.com/mattn/go-sqlite3"
)


func UserExists(user string) error {
    // Открываем базу данных 
    db, err := sql.Open("sqlite3", "users.db")
    if err != nil {
        return err 
    }
    defer db.Close()

    query := "SELECT COUNT(*) FROM users WHERE username = ?"
    var count int
    err = db.QueryRow(query, user).Scan(&count)

    if err != nil {
        return err
    }
    
    if count > 0 {
        return fmt.Errorf("user already exists")
    }

    return nil 
}

func UserLogin(db *sql.DB, userName, password string) error {

    query := "SELECT password FROM users WHERE username =?"
    row := db.QueryRow(query, userName)

    var comparePassword string

    err := row.Scan(&comparePassword)
    if err != nil {
        if err == sql.ErrNoRows {
            return fmt.Errorf("user not found")
        }
    }

    if comparePassword != password {
        return fmt.Errorf("invalid password")
    }

    return nil
}

func UserInsertDB(user User) error {
	// Открываем базу данных 
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err 
	}
	defer db.Close()

	query := "INSERT INTO users (userName, password) VALUES (?,?)"

	_, err = db.Exec(query, user.UserName, user.Password)
	if err != nil {
		return err
	}

	return nil
}