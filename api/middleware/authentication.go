package middleware

import (
	"encoding/base64"
	"fmt"
	"strings"

	"net/http"

	"github.com/access-module/api/db"
	"github.com/access-module/api/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Authentication() gin.HandlerFunc {
	ERROR_MESSAGE := "Unauthorized"
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": ERROR_MESSAGE})
			c.Abort()
			return
		}

		credentialsBase64 := authHeader[len("Basic "):]

		decoded, err := base64.StdEncoding.DecodeString(credentialsBase64)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": ERROR_MESSAGE})
			c.Abort()
			return
		}

		credentials := strings.SplitN(string(decoded), ":", 2)
		if len(credentials) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": ERROR_MESSAGE})
			c.Abort()
			return
		}

		username := credentials[0]
		password := credentials[1]

		db, err := db.Connect()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": ERROR_MESSAGE})
			c.Abort()
			return
		}

		var user model.User
		if err := db.Where("username = ?", username).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": ERROR_MESSAGE})
			c.Abort()
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
		if err != nil {
			fmt.Println("err:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": ERROR_MESSAGE})
			c.Abort()
			return
		}

		c.Next()
	}
}
