package controllers

import (
	"GO_AUTH_JWT/initializers"
	"GO_AUTH_JWT/models"
	"net/http"
	"os"
	"time"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {

	// Get email/passw off req body
	var X struct {
		Email    string `json:"Email"`
		Password string `json:"Password"`
	}

	if err := c.Bind(&X); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body request",
		})
		return
	}

	// Validate email and password
	if strings.TrimSpace(X.Email) == "" || strings.TrimSpace(X.Password) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email and password cannot be empty",
		})
		return
	}

	// hash the password

	hash, err := bcrypt.GenerateFromPassword([]byte(X.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash the password",
		})
		return
	}

	// Create the user

	create := models.User{Email: X.Email, Password: string(hash)}

	result := initializers.DB.Create(&create)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})
		return
	}

	// Reponse
	c.JSON(200, gin.H{})

}

func Login(c *gin.Context) {

	// Récupérer l'email et le mot de passe depuis le corps de la requête
	var X struct {
		Email    string `json:"Email"`
		Password string `json:"Password"`
	}

	if err := c.Bind(&X); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body request",
		})
		return
	}

	// Rechercher l'utilisateur par email

	var user models.User

	initializers.DB.First(&user, "Email=?", X.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalide email or password ",
		})
		return
	}

	// Comparer le mot de passe envoyé avec le hash enregistré

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(X.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalide password ",
		})
		return
	}

	// Récupérer la clé secrète
	secret := os.Getenv("SECRET")
	if secret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "SECRET key is not set or empty",
		})
		return
	}

	// Générer un token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	// Signer le token avec la clé secrète
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	// Retourner le token
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*7, "", "", false, true)

	c.JSON(200, gin.H{
		"token":   tokenString,
		"message": "Connexion réussie",
	})

}

func Validate(c *gin.Context) {

	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
