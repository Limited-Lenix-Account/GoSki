package traffic

import "time"

type UseableTraffic struct {
	ID   string
	COID string
	Name string

	Route     string
	Direction string
	StartMile float64
	EndMile   float64

	TravelTime  int
	UpdatedTime time.Time
}
