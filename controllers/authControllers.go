package controllers

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"
	"github.com/go-playground/validator/v10"

	"backend-golang/config"
	"backend-golang/models"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type RegistrationInput struct {
	Photos   *multipart.FileHeader `form:"image"`
	Name     string `form:"name" binding:"required,alpha"`
	Username string `form:"username" binding:"required,min=5"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

var jwtKey []byte

func init() {
	jwtKey = []byte(os.Getenv("JWT_KEY"))
}

func Register(c *gin.Context) {
	var input RegistrationInput
	log.Printf("JWT Key: %s", jwtKey)

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.Users
	if err := config.DB.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := gonanoid.Generate("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890", 15)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to generate ID"})
		return
	}

	// Proses upload file
	var photoFileName string
	if input.Photos != nil {
		// Generate unique file name with separator
		uniqueCode, _ := gonanoid.Generate("1234567890", 5)
		fileExt := filepath.Ext(input.Photos.Filename)
		photoFileName = fmt.Sprintf("%s-%s%s", id, uniqueCode, fileExt)

		photoPath := filepath.Join("public/images", photoFileName)

		if err := saveFile(input.Photos, photoPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
			return
		}
	}

	users := models.Users{
		ID:       id,
		Name:     input.Name,
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
		Photos:   photoFileName,
	}

	tokenString, err := createJWT(users.ID, "admin")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create JWT"})
		return
	}

	log.Printf("Creating user in database with ID: %s", users.ID)
	if err := config.DB.Create(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "User has been created",
		"data":    users,
		"token":   tokenString,
	})
}

func saveFile(file *multipart.FileHeader, path string) error {
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	openedFile, err := file.Open()
	if err != nil {
		return err
	}
	defer openedFile.Close()

	_, err = io.Copy(out, openedFile)
	if err != nil {
		return err
	}

	return nil
}

func createJWT(userID string, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"exp":  expirationTime.Unix(),
		"sub":  userID,
		"role": role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Printf("Error signing token: %v", err)
		return "", err
	}

	return tokenString, nil
}

type Claims struct {
	jwt.StandardClaims
	Role string `json:"role"`
}

type LoginInput struct {
	Username string `form:"username" binding:"required,min=5"`
	Password string `form:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBind(&input); err != nil {
		log.Printf("Error binding input: %v", err)
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}

	var user models.Users

	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		log.Printf("Error finding user: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		log.Printf("Error comparing password: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	tokenString, err := createJWT(user.ID, "admin")
	if err != nil {
		log.Printf("Error creating JWT: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create JWT"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "User has been logged in",
		"data":    user,
		"token":   tokenString,
	})
}

type CustomersInput struct {
	Name     string `form:"name" binding:"required,alpha"`
	Username string `form:"username" binding:"required,min=5"`
	Password string `form:"password" binding:"required"`
}

func CreateCustomer(c *gin.Context) {
	var input CustomersInput

	if err := c.ShouldBind(&input); err != nil {
		log.Printf("Error binding input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    var existingUser models.Users
	if err := config.DB.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := gonanoid.Generate("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890", 15)
	if err != nil {
		log.Printf("Error generating ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customers := models.Customers{
		ID:       id,
		Name:     input.Name,
		Username: input.Username,
		Password: string(hashedPassword),
	}

	tokenString, err := createJWT(customers.ID, "customer")
	if err != nil {
		log.Printf("Error creating JWT: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create JWT"})
		return
	}

	if err := config.DB.Create(&customers).Error; err != nil {
		log.Printf("Error creating customer in database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Customer has been created",
		"data":    customers,
		"token":   tokenString,
	})
}

func LoginCustomer(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBind(&input); err != nil {
		log.Printf("Error binding input: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var customer models.Customers

	if err := config.DB.Where("username = ?", input.Username).First(&customer).Error; err != nil {
		log.Printf("Error finding customer: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(input.Password)); err != nil {
		log.Printf("Error comparing password: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	tokenString, err := createJWT(customer.ID, "customer")
	if err != nil {
		log.Printf("Error creating JWT: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create JWT"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Customer has been logged in",
		"data":    customer,
		"token":   tokenString,
	})
}
