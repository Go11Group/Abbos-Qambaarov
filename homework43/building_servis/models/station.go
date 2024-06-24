package models

type Stations struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CreateStation struct {
	Name string `json:"name"`
}
