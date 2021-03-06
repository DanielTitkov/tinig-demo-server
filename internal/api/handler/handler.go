package handler

import (
	"github.com/DanielTitkov/tinig-demo-server/internal/app"
	"github.com/DanielTitkov/tinig-demo-server/internal/configs"
	"github.com/DanielTitkov/tinig-demo-server/internal/logger"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Handler struct {
	cfg    configs.Config
	logger *logger.Logger
	app    *app.App
}

func NewHandler(
	e *echo.Echo,
	cfg configs.Config,
	logger *logger.Logger,
	app *app.App,
) *Handler {
	h := &Handler{
		cfg:    cfg,
		logger: logger,
		app:    app,
	}
	h.link(e)
	return h
}

func (h *Handler) link(e *echo.Echo) {
	e.POST("/getToken", h.GetTokenHandler)
	e.POST("/createUser", h.CreateUserHandler)
	e.POST("/getSystemSummary", h.GetSystemSummaryHandler)

	// Restricted group
	restricted := e.Group("/private")
	restricted.Use(middleware.JWT([]byte(h.cfg.Auth.Secret)))
	restricted.POST("/getUser", h.GetUserHandler)
	restricted.POST("/createTask", h.CreateTaskHandler)
	restricted.POST("/createItem", h.CreateItemHandler)
	restricted.POST("/getTasks", h.GetTasksHandler)
	restricted.POST("/updateTask", h.UpdateTaskHandler)
}
