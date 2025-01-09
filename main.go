package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var users = map[string]string{} // Simulasi database pengguna (email -> password)

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
	Data   string `json:"data"`
	Status bool   `json:"status"`
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

	// Save user to "database"
	users[req.Email] = req.Password

	return c.JSON(http.StatusOK, jsonResponse{
		Data:   "User registered successfully",
		Status: true,
	})
}

func Login(c echo.Context) error {
	// Parse request body
	req := new(userRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, jsonResponse{
			Data:   "Invalid request",
			Status: false,
		})
	}

	// Authenticate user
	password, exists := users[req.Email]
	if !exists || password != req.Password {
		return c.JSON(http.StatusUnauthorized, jsonResponse{
			Data:   "Invalid email or password",
			Status: false,
		})
	}

	return c.JSON(http.StatusOK, jsonResponse{
		Data:   "Login successful",
		Status: true,
	})
}
