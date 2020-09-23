package handler

import (
	"net/http"

	"github.com/DanielTitkov/tinig-demo-server/internal/api/model"
	"github.com/DanielTitkov/tinig-demo-server/internal/domain"
	"github.com/labstack/echo"
)

func (h *Handler) CreateItemHandler(c echo.Context) error {
	request := new(model.CreateItemRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	err := h.app.CreateItem(&domain.Item{
		Task:   request.Task,
		Source: request.Source,
		Data:   request.Data,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to create item",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.OKResponse{
		Status:  "ok",
		Message: "item created",
	})
}
