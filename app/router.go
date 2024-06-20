package app

import (
	"net/http"
	"redesign/controller"

	"github.com/julienschmidt/httprouter"
)

type Controllers struct {
	EventController controller.EventController
}

func NewRouter(controllers Controllers) *httprouter.Router {
	router := httprouter.New()

	router.GET("/hello", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		json := `{message: "Hello, World!"}`
		writer.Header().Set("Content-Type", "application/json")
		writer.Write([]byte(json))
	})
	router.GET("/events", controllers.EventController.FindAll)

	return router
}
