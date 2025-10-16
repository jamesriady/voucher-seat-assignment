package domain

type Voucher struct {
	ID           int64        `json:"id"`
	CrewName     string       `json:"crew_name"`
	CrewID       string       `json:"crew_id"`
	FlightNumber string       `json:"flight_number"`
	FlightDate   string       `json:"flight_date"`
	AircraftType AircraftType `json:"aircraft_type"`
	Seats        []string     `json:"seats"`
	CreatedAt    string       `json:"created_at"`
}

type VoucherDB struct {
	ID           int64        `db:"id"`
	CrewName     string       `db:"crew_name"`
	CrewID       string       `db:"crew_id"`
	FlightNumber string       `db:"flight_number"`
	FlightDate   string       `db:"flight_date"`
	AircraftType AircraftType `db:"aircraft_type"`
	Seat1        string       `db:"seat1"`
	Seat2        string       `db:"seat2"`
	Seat3        string       `db:"seat3"`
	CreatedAt    string       `db:"created_at"`
}

type GenerateVouchersRequest struct {
	CrewName     string       `json:"name" validate:"required"`
	CrewID       string       `json:"id" validate:"required"`
	FlightNumber string       `json:"flightNumber" validate:"required"`
	FlightDate   string       `json:"date" validate:"required,datetime=2006-01-02"`
	AircraftType AircraftType `json:"aircraft" validate:"required"`
}

type GenerateVouchersResponse struct {
	Success bool     `json:"success"`
	Error   string   `json:"error,omitempty"`
	Seats   []string `json:"seats,omitempty"`
}

type CheckVoucherRequest struct {
	FlightNumber string `json:"flightNumber" validate:"required"`
	Date         string `json:"date" validate:"required,datetime=2006-01-02"`
}

const MaxVoucherPerFlight = 3
