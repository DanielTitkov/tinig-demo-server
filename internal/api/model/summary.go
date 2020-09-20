package model

import "time"

type (
	GetSystemSummaryResponse struct {
		Users           int       `json:"users"`
		Tasks           int       `json:"tasks"`
		ActiveTasks     int       `json:"activeTasks"`
		Items           int       `json:"items"`
		AvgItemsPerTask float64   `json:"avgItemsPerTask"`
		CreateTime      time.Time `json:"createTime"`
	}
)
