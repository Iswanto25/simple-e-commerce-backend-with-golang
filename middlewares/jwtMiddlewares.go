// middleware/jwtMiddlewares.go

package middlewares

import (
    "net/http"
    "strings"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "backend-golang/controllers"
    // "backend-golang/config"
)

var jwtKey []byte

func init() {
    jwtKey = []byte(os.Getenv("JWT_KEY"))
}

func JWTMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Ambil token dari header Authorization
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        // Format token: Bearer <token>
        tokenString := strings.Split(authHeader, " ")[1]

        // Parse token
        claims := &controllers.Claims{}

        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        // Jika token valid, lanjutkan ke handler berikutnya dengan claims
        c.Set("claims", claims)
        c.Next()
    }
}

func AdminMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        claims, _ := c.Get("claims")
        claim := claims.(*controllers.Claims)

        if claim.Role != "admin" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        c.Next()
    }
}
