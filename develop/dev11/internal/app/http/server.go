package http

import (
	"dev11/internal/models"
	"net/http"
)

type service interface {
	CreateEvent(input models.Event) error
	DeleteEvent(input models.Event) error
	UpdateEvent(input models.Event, oldName string) error
	GetEventForDay(input models.Event) ([]string, error)
	GetEventForWeek(input models.Event) ([]string, error)
	GetEventForMonth(input models.Event) ([]string, error)
}

func NewServer(service service) *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("/create_event", logger(createEvent(service)))
	r.HandleFunc("/update_event", logger(updateEvent(service)))
	r.HandleFunc("/delete_event", logger(deleteEvent(service)))
	r.HandleFunc("/events_for_day", logger(getEventForDay(service)))
	r.HandleFunc("/events_for_week", logger(getEventForWeek(service)))
	r.HandleFunc("/events_for_month", logger(getEventForMonth(service)))

	return r
}
