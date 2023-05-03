package Auth

import (
	"net/http"
	"time"

	// "github.com/99designs/gqlgen/integration/models-go"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/programmer-for-good/flashcardApi/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// type CardsControllerInterface interface {
// 	Register(ctx *gin.Context)
// 	// add other methods you want to implement
// 	Controllers.CardsController
// }

// var t *CardsControllerInterface

func Register(ctx *gin.Context, db *gorm.DB) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the user's password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create a new user with the hashed password
	user.Password = string(hashedPassword)
	if err := db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return a success response
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
	return
}

// Define your JWT claims struct
type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

// Your login endpoint handler
func LoginHandler(c *gin.Context, db *gorm.DB) {
	// Parse the email and password from the request body
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify the email and password
	user, err := GetUserByEmail(loginRequest.Email, db)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Create a JWT token with the user's information and a secret key
	claims := &Claims{
		UserID: int(user.ID),
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // token expires in 24 hours
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret-key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send the JWT token back to the frontend as a response to the login request
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func GetUserByEmail(s string, db *gorm.DB) (*models.User, error) {
	// panic("unimplemented")
	// return s,error()

	// db.Get()
	result := models.User{}
	db.Model(models.User{Email: s}).First(&result)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	return &result, nil
	// SELECT * FROM users WHERE id = 10;

}
