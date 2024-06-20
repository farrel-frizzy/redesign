package controller

import (
	"net/http"
	"redesign/helper"
	"redesign/service"

	"github.com/julienschmidt/httprouter"
)

type EventControllerImpl struct {
	EventService service.EventService
}

func NewEventController(es service.EventService) EventController {
	return &EventControllerImpl{
		EventService: es,
	}
}

func (ec *EventControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	eventResponses := ec.EventService.FindAll(request.Context())
	helper.WriteResponse(writer, eventResponses)
}
