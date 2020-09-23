package app

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"

	"github.com/DanielTitkov/tinig-demo-server/internal/domain"
)

func (a *App) CreateItem(i *domain.Item) error {
	t, err := a.repo.GetTaskByCode(i.Task)
	if err != nil {
		return err
	}

	err = a.ValidateItemData(t.Code, &i.Data)
	if err != nil {
		return err
	}

	dataJSON, err := json.Marshal(i.Data)
	if err != nil {
		return err
	}

	i.Hash = hashFromBytes(dataJSON)
	_, err = a.repo.CreateItem(i, t)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) ValidateItemData(taskCode string, data *domain.ItemData) error {
	task, err := a.repo.GetTaskByCode(taskCode)
	if err != nil {
		return err
	}

	taskType := task.Type
	switch taskType {
	case domain.TaskTypeRandom:
		if data.Value == 0 {
			return errors.New("value is required for task type " + taskType)
		}
	case domain.TaskTypePrice:
		if data.Price == 0 {
			return errors.New("price is required for task type " + taskType)
		}
		if data.PriceMin == 0 {
			return errors.New("priceMin is required for task type " + taskType)
		}
		if data.PriceMax == 0 {
			return errors.New("priceMax is required for task type " + taskType)
		}
	default:
		return errors.New("got unsupported task type: " + taskType)
	}

	return nil
}

func hashFromBytes(b []byte) string { // TODO: maybe move to utils or smth
	hasher := md5.New()
	hasher.Write(b)
	return hex.EncodeToString(hasher.Sum(nil))
}
