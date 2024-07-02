package postgres

import (
	"database/sql"
	models "mymod/module"
)

type StationRepo struct {
	Db *sql.DB
}

func NewStationRepo(db *sql.DB) *StationRepo {
	return &StationRepo{Db: db}
}

func (s *StationRepo) CreateStation(station *models.Stations) error {

	_, err := s.Db.Exec("insert into station(id, name) values ($1, $2)",
		station.Id, station.Name)

	return err
}

func (s *StationRepo) GetStationById(id string) (*models.Stations, error) {
	var station = models.Stations{Id: id}

	err := s.Db.QueryRow("select name from station where id = $1", id).
		Scan(&station.Name)
	if err != nil {
		return nil, err
	}

	return &station, nil
}

func (s *StationRepo) GetStation() (*[]models.Stations, error) {
	rows, err := s.Db.Query("select id, name from station")
    if err!= nil {
        return nil, err
    }
    defer rows.Close()

    var stations []models.Stations
    for rows.Next() {
        var station = models.Stations{}
        err := rows.Scan(&station.Id, &station.Name)
        if err!= nil {
            return nil, err
        }
        stations = append(stations, station)
    }

    return &stations, nil
}

func (s *StationRepo) UpdateStation(station models.Stations,id string) error {
	_, err := s.Db.Exec("update station set name = $1 where id = $2",
        station.Name, id)

    return err
}

func (s *StationRepo) DeleteStation(id string) error {
	_, err := s.Db.Exec("delete from station where id = $1", id)

    return err
}