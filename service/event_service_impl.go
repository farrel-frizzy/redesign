package service

import (
	"context"
	"database/sql"
	"redesign/model/web"
	"redesign/repository"
)

type EventServiceImpl struct {
	EventRepository repository.EventRepository
	DB              *sql.DB
}

func NewEventService(er repository.EventRepository, db *sql.DB) EventService {
	return &EventServiceImpl{EventRepository: er, DB: db}
}

func (es *EventServiceImpl) FindAll(ctx context.Context) []web.EventResponse {
	tx, err := es.DB.Begin()
	if err != nil {
		panic(err)
	}
	events := es.EventRepository.FindAll(ctx, tx)
	var eventResponses []web.EventResponse
	for _, event := range events {
		eventResponses = append(eventResponses, web.EventResponse{
			ID:        event.ID,
			Pamflet:   event.Pamflet,
			Date:      event.Date,
			Title:     event.Title,
			EventType: event.EventType,
		})
	}
	return eventResponses
}
