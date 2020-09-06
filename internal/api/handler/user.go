package handler

import (
	"net/http"

	"github.com/DanielTitkov/tinig-demo-server/internal/domain"
	"github.com/labstack/echo"
)

func (h *Handler) CreateUserHandler(c echo.Context) error {
	user := new(domain.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	err := h.app.CreateUser(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
