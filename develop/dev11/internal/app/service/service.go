package service

import (
	"dev11/internal/models"
	"dev11/internal/repository"
	"fmt"
	"time"
)

//сервис имеет зависимость на абстракциях, место для бизнес логики.

type Service struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateEvent(input models.Event) error {
	err := s.repo.CreateEvent(input)
	if err != nil {
		return fmt.Errorf("[create] event creation error, error: %w", err)
	}
	return nil
}

func (s *Service) UpdateEvent(input models.Event, oldName string) error {
	err := s.repo.UpdateEvent(input, oldName)
	if err != nil {
		return fmt.Errorf("[update] event modification error, error: %w", err)
	}
	return nil
}

func (s *Service) DeleteEvent(input models.Event) error {
	err := s.repo.DeleteEvent(input)
	if err != nil {
		return fmt.Errorf("[delete] error deleting an event, error: %w", err)
	}
	return nil
}

func (s *Service) GetEventForDay(input models.Event) ([]string, error) {
	events, err := s.repo.GetEventForDay(input)
	if err != nil {
		return nil, fmt.Errorf("[getEventForDay] error receiving events, error: %w", err)
	}
	return events, nil
}

func (s *Service) GetEventForWeek(input models.Event) ([]string, error) {
	endDate := getEventsForWeek(input.Date)
	events, err := s.repo.GetEventForWeek(input, endDate)
	if err != nil {
		return nil, fmt.Errorf("[getEventForWeek] error receiving events, error: %w", err)
	}
	return events, nil
}

func (s *Service) GetEventForMonth(input models.Event) ([]string, error) {
	endDate := getEventsForMonth(input.Date)
	events, err := s.repo.GetEventForMonth(input, endDate)
	if err != nil {
		return nil, fmt.Errorf("[getEventForMonth] error receiving events, error: %w", err)
	}
	return events, nil
}

func getEventsForWeek(startDate time.Time) time.Time {
	endDate := startDate.AddDate(0, 0, 7)
	return endDate
}

func getEventsForMonth(startDate time.Time) time.Time {
	endDate := startDate.AddDate(0, 1, 0)
	return endDate
}
