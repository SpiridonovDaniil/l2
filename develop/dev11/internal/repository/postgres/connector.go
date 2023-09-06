package postgres

import (
	"dev11/internal/config"
	"dev11/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Db struct {
	db *sqlx.DB
}

func New(cfg config.Postgres) *Db {
	conn, err := sqlx.Connect("postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.User,
			cfg.Pass,
			cfg.Address,
			cfg.Port,
			cfg.Db,
		))
	if err != nil {
		log.Fatal(err)
	}

	return &Db{db: conn}
}

func (d Db) CreateEvent(input models.Event) error {
	query := "INSERT INTO events (event_name, event_date, user_id) VALUES ($1, $2, $3)"
	_, err := d.db.Exec(query, input.Name, input.Date, input.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (d Db) DeleteEvent(input models.Event) error {
	query := "DELETE FROM events WHERE user_id = $1 AND event_name = $2"
	_, err := d.db.Exec(query, input.UserId, input.Name)
	if err != nil {
		return err
	}

	return nil
}

func (d Db) UpdateEvent(input models.Event, oldName string) error {
	query := "UPDATE events SET event_date = $1, event_name = $2 WHERE user_id = $3 AND event_name = $4"
	_, err := d.db.Exec(query, input.Date, input.Name, input.UserId, oldName)
	if err != nil {
		return err
	}

	return nil
}

func (d Db) GetEventForDay(input models.Event) ([]string, error) {
	query := "SELECT event_name FROM events WHERE event_date = $1 AND user_id = $2"
	rows, err := d.db.Query(query, input.Date, input.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eventNames []string

	for rows.Next() {
		var eventName string
		if err := rows.Scan(&eventName); err != nil {
			return nil, err
		}
		eventNames = append(eventNames, eventName)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return eventNames, nil
}

func (d Db) GetEventForWeek(input models.Event, endDate time.Time) ([]string, error) {
	query := "SELECT event_name FROM events WHERE event_date BETWEEN $1 AND $2 AND user_id = $3"
	rows, err := d.db.Query(query, input.Date, endDate, input.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eventNames []string
	for rows.Next() {
		var eventName string
		err := rows.Scan(&eventName)
		if err != nil {
			return nil, err
		}
		eventNames = append(eventNames, eventName)
	}

	return eventNames, nil
}

func (d Db) GetEventForMonth(input models.Event, endDate time.Time) ([]string, error) {
	query := "SELECT event_name FROM events WHERE event_date BETWEEN $1 AND $2 AND user_id = $3"
	rows, err := d.db.Query(query, input.Date, endDate, input.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eventNames []string
	for rows.Next() {
		var eventName string
		err := rows.Scan(&eventName)
		if err != nil {
			return nil, err
		}
		eventNames = append(eventNames, eventName)
	}

	return eventNames, nil
}
