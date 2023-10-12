package middleware

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"net/http"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		credentialsBase64 := authHeader[len("Basic "):]

		decoded, err := base64.StdEncoding.DecodeString(credentialsBase64)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		credentials := strings.SplitN(string(decoded), ":", 2)
		if len(credentials) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials format"})
			c.Abort()
			return
		}

		// username := credentials[0]
		// password := credentials[1]
		// fmt.Println("👉 username:", username)
		// fmt.Println("👉 password:", password)
		fmt.Println("👉 Authentication: ✅")

		// Go to db and check for the presence of the user:

		c.Next()
	}
}
