package handler

import (
	"net/http"

	"github.com/DanielTitkov/tinig-demo-server/internal/api/model"
	"github.com/DanielTitkov/tinig-demo-server/internal/api/util"
	"github.com/DanielTitkov/tinig-demo-server/internal/domain"
	"github.com/labstack/echo"
)

func (h *Handler) CreateTaskHandler(c echo.Context) error {
	request := new(model.CreateTaskRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	username, err := util.UsernameFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "token is invalid",
			Error:   err.Error(), // TODO: bind with echo logging
		})
	}

	err = h.app.CreateTask(&domain.Task{
		User:        username,
		Type:        request.Type,
		Slug:        request.Slug,
		Title:       request.Title,
		Description: request.Description,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to create task",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.OKResponse{
		Status:  "ok",
		Message: "task created",
	})
}
