package middleware

import (
	"GO_AUTH_JWT/initializers"
	"GO_AUTH_JWT/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// Obtenir le token du header Authorization
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		// Essayer de récupérer depuis le cookie comme fallback
		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		authHeader = "Bearer " + tokenString
	}

	// Extraire le token du format "Bearer <token>"
	tokenString := authHeader
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		tokenString = authHeader[7:]
	}

	// Décoder/valider le token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("Unexpected signing method:", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secret := os.Getenv("SECRET")
		return []byte(secret), nil
	})

	if err != nil {
		log.Println("Erreur de parsing du token:", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Vérifier l'expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			log.Println("Token expiré")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Trouver l'utilisateur avec le sub du token
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			log.Println("Utilisateur non trouvé")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Attacher à la requête
		c.Set("user", user)

		// Continuer
		c.Next()
	} else {
		log.Println("Claims invalides")
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
