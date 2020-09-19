package app

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"

	"github.com/DanielTitkov/tinig-demo-server/internal/domain"
)

func (a *App) CreateItem(i *domain.Item) error {
	t, err := a.repo.GetTaskByCode(i.Task)
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

func hashFromBytes(b []byte) string {
	hasher := md5.New()
	hasher.Write(b)
	return hex.EncodeToString(hasher.Sum(nil))
}
