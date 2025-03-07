package users 

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	ID int 			`json:"id"`
	jwt.StandardClaims
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var key = []byte("wj2h4kn1h4wau8am3n662Nmad")


// Функция для создания JWT токена
func CreateToken(username string, id int) (string, error) {

    tokenLifeTime := time.Now().Add(time.Hour * 1)

    claims := &Claims{
        Username: 		username,
		ID: 	  		id,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt:  tokenLifeTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // подписываем токен
    signedToken, err := token.SignedString(key)
    if err != nil {
        return "", err
    }

    return signedToken, nil
}
