package middleware

import (
	"net/http"
	"my-go-project/users"

	"github.com/gin-gonic/gin"
)

func ValideUserData() gin.HandlerFunc {
    return func (c *gin.Context) {
        var user users.User

        if err := c.ShouldBindBodyWithJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.JSON(http.StatusBadRequest, gin.H{"error": "JSON request error"})
            c.Abort()
            return
        }
        c.Set("Validated user: ", user)
        c.Next()
    }
}
