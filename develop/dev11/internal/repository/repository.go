package repository

import (
	"dev11/internal/models"
	"time"
)

type Repository interface {
	CreateEvent(input models.Event) error
	DeleteEvent(input models.Event) error
	UpdateEvent(input models.Event, oldName string) error
	GetEventForDay(input models.Event) ([]string, error)
	GetEventForWeek(input models.Event, endDate time.Time) ([]string, error)
	GetEventForMonth(input models.Event, endDate time.Time) ([]string, error)
}
