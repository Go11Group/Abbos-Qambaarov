package handler

import (
	"database/sql"
	"mymod/storage/postgres"
)

type handler struct {
	User *postgres.UserRepo
	Card *postgres.CardRepo
	Station *postgres.StationRepo
	Terminal *postgres.TerminalRepo
}

func NewHandler(db *sql.DB) *handler {
	return &handler{
		User: postgres.NewUserRepo(db),
	}
}

func NewCardRepo(db *sql.DB) *handler {
	return &handler{
		Card: postgres.NewCardRepo(db),
	}
}

func NewStation(db *sql.DB) *handler {
	return &handler{
        Station: postgres.NewStationRepo(db),
    }
}

func NewTerminal(db *sql.DB) *handler {
	return &handler{
        Terminal: postgres.NewTerminalRepo(db),
    }
}