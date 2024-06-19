package domain

import "time"

type Event struct {
	ID          int
	Title       string
	Description string
	EventDate   time.Time
	Location    string
	Image       string
}
