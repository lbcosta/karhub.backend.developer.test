package domain

type Beer struct {
	ID             int      `json:"id"`
	Style          *string  `json:"style"`
	MinTemperature *float64 `json:"min_temperature"`
	MaxTemperature *float64 `json:"max_temperature"`
}
