package service

import (
	"context"
	"redesign/model/web"
)

type EventService interface {
	//Create(ctx context.Context, request web.CreateEventRequest) web.EventResponse
	//Update(ctx context.Context, request web.UpdateEventRequest) web.EventResponse
	//Delete(ctx context.Context, eventId int) web.WebResponse
	//FindById(ctx context.Context, eventId int) web.EventResponse
	FindAll(ctx context.Context) []web.EventResponse
}
