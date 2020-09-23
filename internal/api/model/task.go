package model

import (
	"time"

	"github.com/DanielTitkov/tinig-demo-server/internal/domain"
)

type (
	CreateTaskRequest struct {
		Title       string            `json:"title"`
		Type        string            `json:"type"`
		Slug        string            `json:"slug"`
		Description string            `json:"description"`
		Params      domain.TaskParams `json:"params"`
		Meta        domain.TaskMeta   `json:"meta"`
	}
	UpdateTaskRequest struct {
		Code        string            `json:"code"` // cannot be changed, used for task finding
		Title       string            `json:"title"`
		Description string            `json:"description"`
		Active      bool              `json:"active"`
		Params      domain.TaskParams `json:"params"`
		Meta        domain.TaskMeta   `json:"meta"`
	}
	GetTasksRequest struct {
		WithItems   bool `json:"withItems"`
		Deactivated bool `json:"deactivated"`
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
		Active      bool                   `json:"active"`
		Params      domain.TaskParams      `json:"params,omitempty"`
		Meta        domain.TaskMeta        `json:"meta,omitempty"`
		Items       []GetTasksResponseItem `json:"items,omitempty"`
	}
	GetTasksResponseItem struct {
		Source     string          `json:"source"`
		CreateTime time.Time       `json:"createTime"`
		Data       domain.ItemData `json:"data"`
	}
)
