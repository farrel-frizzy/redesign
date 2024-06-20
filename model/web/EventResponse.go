package web

import "time"

type EventResponse struct {
	ID        int       `json:"id"`
	Pamflet   string    `json:"pamflet"`
	Date      time.Time `json:"date"`
	Title     string    `json:"title"`
	EventType string    `json:"eventType"`
}
