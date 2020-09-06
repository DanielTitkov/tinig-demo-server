package handler

import (
	"github.com/DanielTitkov/tinig-demo-server/internal/configs"
	"github.com/DanielTitkov/tinig-demo-server/internal/logger"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Handler struct {
	cfg    configs.Config
	logger *logger.Logger
}

func InitHandler(
	e *echo.Echo,
	cfg configs.Config,
	logger *logger.Logger,
) *Handler {
	h := &Handler{
		cfg:    cfg,
		logger: logger,
	}
	h.link(e)
	return h
}

func (h *Handler) link(e *echo.Echo) {
	e.POST("/login", h.LoginHandler)

	// Restricted group
	restricted := e.Group("/restricted")
	restricted.Use(middleware.JWT([]byte("secret")))
}
