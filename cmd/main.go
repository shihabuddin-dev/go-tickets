package main

import (
	"go-tickets/internal/config"
	"go-tickets/internal/user"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}

func main() {
	config := config.LoadEnv()
	dsn := config.Dsn

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		panic("failed to connect database")
	} else {
		println("Database connection successful")
	}
	// Auto-migrate the User model to create the users table if it doesn't exist
	db.AutoMigrate(&user.User{})

	e := echo.New()
	e.Use(middleware.RequestLogger())

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Validator = &CustomValidator{validator: validator.New()}

	// user route registration
	user.RegisterRoutes(e, db)

	if err := e.Start(":" + config.Port + ""); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
