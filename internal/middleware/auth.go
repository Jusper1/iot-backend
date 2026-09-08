package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Authorization token diperlukan",
			})
			c.Abort()
			return
		}
	parts := strings.SplitN(authHeader, " ", 2)

	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Format token tidak valid",
		})
		c.Abort()
		return
	}

	tokenString := parts[1]

	secret := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(
		tokenString, 
		func(token *jwt.Token) (interface{}, error) {
		if _, ok := 
		token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, 
			jwt.ErrSignatureInvalid
		}
		
		return []byte(secret), nil
	})	

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Token tidak valid atau expired",
		})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Token claim tidak valid",
		})
		c.Abort()
		return
	}

	c.Set("claims", claims)

	c.Next()
	}
}