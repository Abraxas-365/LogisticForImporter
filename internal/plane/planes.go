package plane

import "time"

type Plane struct {
	ID            int       `json:"id"`
	Departure     string    `json:"departure"`
	Destination   string    `json:"destination"`
	DepartureDate time.Time `json:"departure_date"`
	ArrivalDate   time.Time `json:"arrival_date"`
}
