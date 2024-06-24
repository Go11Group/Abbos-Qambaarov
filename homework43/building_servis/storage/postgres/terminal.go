package postgres

import (
	"database/sql"
	"mymod/models"
	"github.com/google/uuid"
)

type TerminalRepo struct {
	Db *sql.DB
}

func NewTerminalRepo(db *sql.DB) *TerminalRepo {
	return &TerminalRepo{Db: db}
}

func (s *TerminalRepo) CreateTerminal(terminal *models.CreateTerminal) error {

	_, err := s.Db.Exec("insert into terminal(id, name) values ($1, $2)",
		uuid.NewString(), terminal.Stationid)

	return err
}

func (s *TerminalRepo) GetTerminalById(id string) (*models.Terminals, error) {
	var terminal = models.Terminals{Id: id}

	err := s.Db.QueryRow("select name from terminal where id = $1", id).
		Scan(&terminal.Stationid)
	if err != nil {
		return nil, err
	}

	return &terminal, nil
}

func (s *TerminalRepo) GetTerminals() (*[]models.Terminals, error) {
	rows, err := s.Db.Query("select id, name from terminal")
    if err!= nil {
        return nil, err
    }
    defer rows.Close()

    var terminals []models.Terminals
    for rows.Next() {
        var terminal = models.Terminals{}
        err := rows.Scan(&terminal.Id, &terminal.Stationid)
        if err!= nil {
            return nil, err
        }
        terminals = append(terminals, terminal)
    }

    return &terminals, nil
}

func (s *TerminalRepo) UpdateTerminal(terminal models.CreateTerminal ,id string) error {
	_, err := s.Db.Exec("update terminal set name = $1 where id = $2",
        terminal.Stationid, id)

    return err
}

func (s *TerminalRepo) DeleteTerminal(id string) error {
	_, err := s.Db.Exec("delete from terminal where id = $1", id)

    return err
}
