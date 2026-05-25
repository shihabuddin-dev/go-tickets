package main

import (
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

// User model
type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required" gorm:"type:varchar(100)"`
	Email    string `json:"email" validate:"required,email" gorm:"type:varchar(100); uniqueIndex; not null"`
	Password string `json:"password" validate:"required,min=6,max=100" gorm:"type:varchar(100) not null"`
}

func main() {
	dsn := "DATABASE_URL"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	db.AutoMigrate(&User{}) // Auto-migrate the User model to create the users table if it doesn't exist
	if err != nil {
		panic("failed to connect database")
	} else {
		println("Database connection successful")
	}

	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users", func(c *echo.Context) error {
		return c.String(http.StatusOK, "This is From Users Routes")
	})

	e.POST("/users", func(c *echo.Context) (err error) {
		newUser := new(User)
		if err = c.Bind(newUser); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
		if err = c.Validate(newUser); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{"error": err.Error()})
		}

		// save to database
		result := db.Create(newUser)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{"error": result.Error.Error()})
		}

		return c.JSON(http.StatusOK, newUser)
	})

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
