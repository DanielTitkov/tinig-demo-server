package domain

import "time"

// task types
const (
	TaskTypeRandom = "random"
	TaskTypePrice  = "price"
)

type (
	// User holds user data
	User struct {
		ID           int  // id is passed to domain model for simplicity
		Service      bool // if user is a service // TODO: only service users create items
		Username     string
		Password     string
		PasswordHash string
		Email        string // TODO: add ent validation
	}
	// Task represents data backend task
	Task struct {
		ID          int    // id is passed to domain model for simplicity
		Code        string // unique task code for external systems
		User        string // username
		Type        string // task type code
		Slug        string
		Title       string
		Description string
		Active      bool
		Items       []Item
		Meta        TaskMeta
		Params      TaskParams
	}
	// TaskMeta from something
	TaskMeta struct{}
	// TaskParams holds task running params provided by user
	TaskParams struct {
		Random TaskParamsRandom `json:"random,omitempty"`
		Price  TaskParamsPrice  `json:"price,omitempty"`
	}
	// TaskParamsRandom hold params from 'random' task
	TaskParamsRandom struct {
		Min int `json:"min"`
		Max int `json:"max"`
	}
	// TaskParamsPrice hold params from 'price' task
	TaskParamsPrice struct{}
	// TaskType holds task types e.g. 'price', 'random'
	TaskType struct {
		ID   int
		Code string
	}
	// Item in an piece of data collected by a task during one run
	Item struct {
		ID         int
		Hash       string
		Source     string
		Task       string // task code (from external system)
		Data       ItemData
		CreateTime time.Time
	}
	// ItemData holds item payload
	ItemData struct {
		// fot task type random
		Value int `json:"value"`
		// for task type price
		Price    float64 `json:"price,omitempty"`
		PriceMax float64 `json:"priceMax,omitempty"`
		PriceMin float64 `json:"priceMin,omitempty"`
		Product  string  `json:"product,omitempty"`
		Seller   string  `json:"seller,omitempty"`
		// for task type word
		Word string  `json:"word,omitempty"`
		Rate float64 `json:"rate,omitempty"`
		// common
		Platform string `json:"platform,omitempty"`
	}
	// SystemSymmary holds system state gathered by the relevant job
	SystemSymmary struct {
		ID              int
		Users           int
		Tasks           int
		ActiveTasks     int
		Items           int
		AvgItemsPerTask float64
		CreateTime      time.Time
	}
)
