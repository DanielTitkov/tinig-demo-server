package handler

import (
	"net/http"

	"github.com/DanielTitkov/tinig-demo-server/internal/domain"
	"github.com/labstack/echo"
)

func (h *Handler) GetTokenHandler(c echo.Context) error {
	u := new(domain.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	valid, err := h.app.ValidateUserPassword(u)
	if err != nil {
		return err
	}
	if !valid {
		return echo.ErrUnauthorized
	}

	token, err := h.app.GetUserToken(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
