package alerts

type UseableAlert struct {
	ID          string
	Reason      string
	Tooltip     string
	Description string

	Route          string
	Direction      string
	StartMile      float64
	EndMile        float64
	PrimaryPoint   Point
	SecondaryPoint Point

	SingleMile  float64
	SinlgePoint Point

	BeginTime  int64
	EndTime    int64
	UpdateTime int64
}

type Point struct {
	Lat             float64
	Long            float64
	LinearReference float64
}
