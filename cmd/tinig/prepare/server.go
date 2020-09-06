package prepare

import (
	"github.com/DanielTitkov/tinig-demo-server/internal/api/handler"
	"github.com/DanielTitkov/tinig-demo-server/internal/configs"
	"github.com/DanielTitkov/tinig-demo-server/internal/logger"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewServer(cfg configs.Config, logger *logger.Logger) *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	handler.InitHandler(e, cfg, logger)
	return e
}
