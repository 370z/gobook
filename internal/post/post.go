package post

import (
	"gobook/internal/utils"
	"net/http"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId uint
}

type Service struct {
	DB *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{DB: db}
}

func (s *Service) Create(c echo.Context) error {
	p := new(Post)
	if err := c.Bind(p); err != nil {
		return err
	}

	//Get userId
		// Get jwtCustomClaims utils
	user := utils.GetJWTClaims(c)

	p.UserId = user.ID
	if err := s.DB.Create(&p).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Something wrong")
	}
	// Response
	return c.JSON(http.StatusCreated, echo.Map{
		"data":p,
		"message": "Post created successfully",
	})

}
