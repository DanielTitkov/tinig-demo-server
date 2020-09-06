package handler

import (
	"net/http"
	"time"

	"github.com/DanielTitkov/tinig-demo-server/internal/api/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func (h *Handler) LoginHandler(c echo.Context) error {

	u := new(model.LoginModel)
	if err := c.Bind(u); err != nil {
		return err
	}

	// Throws unauthorized error
	if u.Username != "jon" || u.Password != "foo" {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = u.Username
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
