package handler

import (
	"net/http"

	"github.com/DanielTitkov/tinig-demo-server/internal/api/model"
	"github.com/labstack/echo"
)

func (h *Handler) GetSystemSummaryHandler(c echo.Context) error {
	sum, err := h.app.GetSystemSummary()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to get system summary",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.GetSystemSummaryResponse{
		Users:           sum.Users,
		Tasks:           sum.Tasks,
		ActiveTasks:     sum.ActiveTasks,
		Items:           sum.Items,
		AvgItemsPerTask: sum.AvgItemsPerTask,
		CreateTime:      sum.CreateTime,
	})
}
