package app

import (
	"github.com/DanielTitkov/tinig-demo-server/internal/domain"
)

func (a *App) CreateSystemSummary() error {
	userCount, err := a.repo.GetUserCount()
	if err != nil {
		return err
	}

	totalTaskCount, err := a.repo.GetTaskCount(false)
	if err != nil {
		return err
	}

	activeTaskCount, err := a.repo.GetTaskCount(true)
	if err != nil {
		return err
	}

	itemCount, err := a.repo.GetItemCount()
	if err != nil {
		return err
	}

	var avgItemsPerTask float64 = float64(itemCount) / float64(totalTaskCount)
	_, err = a.repo.CreateSystemSummary(&domain.SystemSymmary{
		Users:           userCount,
		Tasks:           totalTaskCount,
		ActiveTasks:     activeTaskCount,
		Items:           itemCount,
		AvgItemsPerTask: avgItemsPerTask,
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *App) GetSystemSummary() (*domain.SystemSymmary, error) {
	return a.repo.GetLatestSystemSummary()
}
