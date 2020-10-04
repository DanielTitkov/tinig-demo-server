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
		Params:      request.Params,
		Meta:        request.Meta,
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

func (h *Handler) UpdateTaskHandler(c echo.Context) error {
	request := new(model.UpdateTaskRequest)
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

	err = h.app.UpdateTask(&domain.Task{
		User:        username,     // immutable
		Code:        request.Code, // immutable
		Active:      request.Active,
		Title:       request.Title,
		Description: request.Description,
		Meta:        request.Meta,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to update task",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.OKResponse{
		Status:  "ok",
		Message: "task updated",
	})
}

func (h *Handler) GetTasksHandler(c echo.Context) error {
	request := new(model.GetTasksRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	username, err := util.UsernameFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "token is invalid",
			Error:   err.Error(),
		})
	}

	tasks, err := h.app.GetTasks(&domain.User{Username: username}, request.WithItems)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to get user tasks",
			Error:   err.Error(),
		})
	}

	var resp model.GetTasksResponse
	for _, t := range tasks {
		var items []model.GetTasksResponseItem
		for _, i := range t.Items {
			items = append(items, model.GetTasksResponseItem{
				Source:     i.Source,
				CreateTime: i.CreateTime,
				Data:       i.Data,
			})
		}

		resp.Tasks = append(resp.Tasks, model.GetTasksResponseTask{
			Code:        t.Code,
			Type:        t.Type,
			Slug:        t.Slug,
			Title:       t.Title,
			Active:      t.Active,
			Description: t.Description,
			Params:      t.Params,
			Meta:        t.Meta,
			Items:       items,
		})
	}

	return c.JSON(http.StatusOK, resp)
}
