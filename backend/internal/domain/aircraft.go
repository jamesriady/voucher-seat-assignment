package domain

// AircraftType represents the type of aircraft for a flight.
type AircraftType string

const (
	ATR          AircraftType = "ATR"
	Airbus320    AircraftType = "Airbus 320"
	Boeing737Max AircraftType = "Boeing 737 Max"
)

const (
	ATRMaxRows          = 18
	Airbus320MaxRows    = 32
	Boeing737MaxMaxRows = 32
)

var (
	ATRSeatsPerRow          = []string{"A", "C", "D", "F"}
	Airbus320SeatsPerRow    = []string{"A", "B", "C", "D", "E", "F"}
	Boeing737MaxSeatsPerRow = []string{"A", "B", "C", "D", "E", "F"}
)
