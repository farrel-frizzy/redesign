package repository

import (
	"context"
	"database/sql"
	"redesign/model/domain"
)

type EventRepository interface {
	//Save(ctx context.Context, tx *sql.Tx, event domain.Event) domain.Event
	//Update(ctx context.Context, tx *sql.Tx, event domain.Event) domain.Event
	//Delete(ctx context.Context, tx *sql.Tx, eventId int)
	//FindById(ctx context.Context, tx *sql.Tx, eventId int) (*domain.Event, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Event
}
