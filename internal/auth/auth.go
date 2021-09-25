package auth

import (
	"errors"
	"fmt"
	"gobook/internal/user"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type JwtCustomClaims struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

type Service struct {
	DB *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{DB: db}
}

func (s *Service) Login(c echo.Context) error {
	email := c.FormValue("email")

	var u user.User
	if err := s.DB.First(&u, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// return c.JSON(http.StatusUnauthorized,echo.Map{
			// 	"message" : "Not Found",
			// })
			return echo.ErrUnauthorized
		}
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Something bad happened",
		})
	}
	//Payload
	claims := &JwtCustomClaims{
		u.ID,
		u.Name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	//Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return err
	}
	fmt.Println(err)
	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

