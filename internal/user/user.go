package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// User
type User struct {
	// ID       int    `json:"id" xml:"id"`
	// Username string `json:"username" xml:"username"`
	// Password string `json:"password" xml:"password"`
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int8   `json:"age"`
	IsActive bool   `json:"is_active"`
}
type Service struct {
	DB *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{DB: db}
}

var users []User

func (s *Service) CreateUser(c echo.Context) error {

	name := c.FormValue("name")
	email := c.FormValue("email")
	age := c.FormValue("age")
	isActive := c.FormValue("is_active")
	fmt.Println(users)
	if findName := s.DB.Where(&User{Name: name}, "name").Find(&users); findName.RowsAffected > 0 {
		fmt.Println(findName.RowsAffected)
		return echo.NewHTTPError(400, "Name exist")
	}
	if emailExist := s.DB.Where(&User{Email: email}, "email").Find(&users); emailExist.RowsAffected > 0 {
		return echo.NewHTTPError(400, "Email exist")
	}

	convAge, err := strconv.ParseInt(age, 10, 8)
	if err != nil {
		return echo.NewHTTPError(400, "Cannot convert age string to int")
	}
	convIsActive, err := strconv.ParseBool(isActive)
	if err != nil {
		return echo.NewHTTPError(400, "Cannot convert is_active to bool")
	}

	// var user = User{}
	u := User{
		Name:     name,
		Email:    email,
		Age:      int8(convAge),
		IsActive: bool(convIsActive),
	}

	if createErr := s.DB.Create(&u).Error; createErr != nil {
		return echo.NewHTTPError(500, "Failed to insert data to database")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "User Created",
		"success": true,
	})
}

func (s *Service) GetUsers(c echo.Context) error {

	if err := s.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(500, "Failed to get User")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"users":   users,
		"success": true,
	})
}

func (s *Service) GetUser(c echo.Context) error {

	id := c.Param("id")
	user := User{}

	result := s.DB.First(&user, id)

	if result.RowsAffected == 0 {
		return echo.NewHTTPError(404, "User not found")
	}

	if err := result.Error; err != nil {
		return echo.NewHTTPError(500, "Failed to get User")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user":    user,
		"success": true,
	})
}

func (s *Service) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	u := new(User)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(400, "Bad request")
	}
	if err := s.DB.First(&u, id).Error; err != nil {
		return echo.NewHTTPError(404, "User not found")
	}

	if err := s.DB.Where(id).Updates(u).Error; err != nil {
		return echo.NewHTTPError(500, "Failed to update data to database")
	}
	return c.JSON(200, echo.Map{
		"message": "User Updated",
		"success": true,
	})
}
func (s *Service) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	user := User{}

	if err := s.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(404, "User not found")
	}

	if err := s.DB.Delete(&user, id).Error; err != nil {
		return echo.NewHTTPError(500, "Server error")
	}

	return c.JSON(200, echo.Map{
		"message": "User Deleted",
		"success": true,
	})
}
