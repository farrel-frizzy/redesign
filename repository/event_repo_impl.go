package repository

import (
	"context"
	"database/sql"
	"fmt"
	"redesign/model/domain"
	"time"
)

type EventRepositoryImpl struct {
}

func NewEventRepository() EventRepository {
	return &EventRepositoryImpl{}
}

func (er *EventRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Event {
	Query := `SELECT id, pamflet, title, date, eventType FROM events`
	rows, err := tx.QueryContext(ctx, Query)
	if err != nil {
		panic(err)
	}

	events := []domain.Event{}
	for rows.Next() {
		event := domain.Event{}
		var eventDate []byte
		err := rows.Scan(&event.ID, &event.Pamflet, &event.Title, &eventDate, &event.EventType)
		if err != nil {
			panic(err)
		}

		event.Date, err = time.Parse("2006-01-02", string(eventDate))
		if err != nil {
			panic(err)
		}

		events = append(events, event)
		fmt.Println(events)
	}

	return events
}
