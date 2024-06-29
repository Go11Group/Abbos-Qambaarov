package postgres

import (
	"database/sql"
	models "mymod/module"
)

type CardRepo struct {
	Db *sql.DB
}

func NewCardRepo(db *sql.DB) *CardRepo {
	return &CardRepo{Db: db}
}

func (s *CardRepo) CreateCard(Card *models.Cards) error {

	_, err := s.Db.Exec("insert into cards(id, numbar, user_id) values ($1, $2, $3)",
		Card.Id, Card.Number, Card.UserId)

	return err
}

func (s *CardRepo) GetCardById(id string) (*models.Cards, error) {
	var card = models.Cards{Id: id}

	err := s.Db.QueryRow("select name from cards where id = $1", id).
		Scan(&card.Number, &card.UserId)
	if err != nil {
		return nil, err
	}

	return &card, nil
}

func (s *CardRepo) GetCard() (*[]models.Cards, error) {
	rows, err := s.Db.Query("select id, numbar, user_id from cards")
    if err!= nil {
        return nil, err
    }
    defer rows.Close()

    var cards []models.Cards
    for rows.Next() {
        var card = models.Cards{}
        err := rows.Scan(&card.Id, &card.Number, &card.UserId)
        if err!= nil {
            return nil, err
        }
        cards = append(cards, card)
    }

    return &cards, nil
}

func (s *CardRepo) UpdateCard(Card models.Cards, id string) error {
	_, err := s.Db.Exec("update cards set name = $1 where id = $2",
        Card.Number, Card.UserId, id)

    return err
}

func (s *CardRepo) DeleteCard(id string) error {
	_, err := s.Db.Exec("delete from cards where id = $1", id)

    return err
}