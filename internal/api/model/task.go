package model

import "github.com/DanielTitkov/tinig-demo-server/internal/domain"

type (
	CreateTaskRequest struct {
		Title       string `json:"title"`
		Type        string `json:"type"`
		Slug        string `json:"slug"`
		Description string `json:"description"`
	}
	GetTasksResponse struct {
		Tasks []GetTasksResponseTask `json:"tasks"`
	}
	GetTasksResponseTask struct {
		// User        string
		Code        string                 `json:"code"`
		Type        string                 `json:"type"`
		Slug        string                 `json:"slug"`
		Title       string                 `json:"title"`
		Description string                 `json:"description"`
		Items       []GetTasksResponseItem `json:"items,omitempty"`
	}
	GetTasksResponseItem struct {
		Source string          `json:"source"`
		Data   domain.ItemData `json:"data"`
	}
)
