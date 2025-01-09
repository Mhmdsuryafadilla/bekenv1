package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt" // For password hashing
	"time" // For generating IDs
	"strconv"
	"math/rand"
)

// Define a struct to represent a user.
type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

var users = map[string]User{} // Use map[string]User for user details.

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", Welcome)
	e.GET("/get-user", GetUser)
	e.GET("/get-order", GetOrder)
	e.GET("/get-product", GetProduct)
	e.GET("/get-city", GetCity)
	e.POST("/register", Register)
	e.POST("/login", Login)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

type jsonResponse struct {
	Data   interface{} `json:"data"`
	Status bool        `json:"status"`
	User   *User       `json:"user,omitempty"`
}

// Handler
func Welcome(c echo.Context) error {
	welcome := fmt.Sprintln("Welcome To Website Test API \n 1. /get-user \n 2. /get-order \n 3. /get-product \n 4. /register \n 5. /login")
	return c.String(http.StatusOK, welcome)
}

func GetUser(c echo.Context) error {
	response := jsonResponse{
		Data:   "Data User Berhasil di Get",
		Status: true,
	}
	return c.JSON(http.StatusOK, response)
}

func GetOrder(c echo.Context) error {
	response := jsonResponse{
		Data:   "Data Order Berhasil di Get",
		Status: true,
	}
	return c.JSON(http.StatusOK, response)
}

func GetProduct(c echo.Context) error {
	response := jsonResponse{
		Data:   "Data Product Berhasil di Get",
		Status: true,
	}
	return c.JSON(http.StatusOK, response)
}

func GetCity(c echo.Context) error {
	response := jsonResponse{
		Data:   "Data City Berhasil di Get",
		Status: true,
	}
	return c.JSON(http.StatusOK, response)
}

type userRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c echo.Context) error {
	// Parse request body
	req := new(userRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, jsonResponse{
			Data:   "Invalid request",
			Status: false,
		})
	}

	// Check if user already exists
	if _, exists := users[req.Email]; exists {
		return c.JSON(http.StatusConflict, jsonResponse{
			Data:   "User already exists",
			Status: false,
		})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, jsonResponse{
			Data:   "Error hashing password",
			Status: false,
		})
	}

	// Generate a simple ID using current timestamp
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(10000)
	timestamp := time.Now().Unix()
	strTimestamp := strconv.FormatInt(timestamp, 10)
	strRandomInt := strconv.Itoa(randomInt)

	id, err := strconv.ParseInt(strTimestamp+strRandomInt, 10, 64)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, jsonResponse{
			Data:   "Error generating user id",
			Status: false,
		})
	}

	// Save user to "database"
	user := User{
		ID:       id,
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	users[req.Email] = user

	return c.JSON(http.StatusOK, jsonResponse{
		Data:   "User registered successfully",
		Status: true,
		User:   &user,
	})
}

func Login(c echo.Context) error {
	// Parse request body
	req := new(loginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, jsonResponse{
			Data:   "Invalid request",
			Status: false,
		})
	}

	// Authenticate user
	user, exists := users[req.Email]
	if !exists {
		return c.JSON(http.StatusUnauthorized, jsonResponse{
			Data:   "Invalid email or password",
			Status: false,
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, jsonResponse{
			Data:   "Invalid email or password",
			Status: false,
		})
	}

	return c.JSON(http.StatusOK, jsonResponse{
		Data:   "Login successful",
		Status: true,
		User:   &user,
	})
}