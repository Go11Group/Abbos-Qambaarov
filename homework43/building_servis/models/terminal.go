package models

type Terminals struct {
	Id   string `json:"id"`
	Stationid string `json:"station_id"`
}

type CreateTerminal struct {
	Stationid string `json:"station_id"`
}
