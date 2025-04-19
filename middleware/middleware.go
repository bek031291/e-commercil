package middleware

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}
		client := gocloak.NewClient("http://localhost:8280")
		token := authHeader
		rptResult, err := client.RetrospectToken(c, token, "admin-cli", "KG1yDiz8Gp1d3CWWBJ4HxprPsE2GsCnP", "master")
		if err != nil || rptResult == nil || !*rptResult.Active {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Next()
	}
}
