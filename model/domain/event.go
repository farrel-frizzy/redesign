package domain

import "time"

type Event struct {
	ID        int
	Pamflet   string
	Date      time.Time
	Title     string
	EventType string
}
