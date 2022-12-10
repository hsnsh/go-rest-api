package models

import "time"

type ProductDto struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	CreationTime time.Time `json:"creationtime"`
	UpdateTime   time.Time `json:"updatetime"`
}

type ProductCreateDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProductUpdateDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
