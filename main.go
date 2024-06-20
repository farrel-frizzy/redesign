package main

import (
	"log"
	"net/http"
	"redesign/app"
	"redesign/controller"
	"redesign/middleware"
	"redesign/repository"
	"redesign/service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	eventRepository := repository.NewEventRepository()
	eventService := service.NewEventService(eventRepository, db)
	EventController := controller.NewEventController(eventService)
	controllers := app.Controllers{EventController: EventController}
	router := app.NewRouter(controllers)

	middlewareCors := middleware.CORS(router)

	// Start the server
	err := http.ListenAndServe(":8080", middlewareCors)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	log.Println("Server is running on port 8080")

}
