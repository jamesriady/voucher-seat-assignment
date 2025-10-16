package repository

import (
	"log"
	"time"

	"backend/internal/domain"

	"github.com/jmoiron/sqlx"
)

type VoucherRepository interface {
	Migrate() error
	Exists(flightNumber, flightDate string) (bool, error)
	Create(voucher domain.Voucher) error
}

type sqliteRepository struct {
	db *sqlx.DB
}

// NewVoucherRepository creates a new SQLite-based voucher repository.
func NewVoucherRepository(db *sqlx.DB) VoucherRepository {
	return &sqliteRepository{db: db}
}

// Migrate creates the necessary database schema.
func (r *sqliteRepository) Migrate() error {
	query := `
    CREATE TABLE IF NOT EXISTS vouchers (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        crew_name TEXT NOT NULL,
        crew_id TEXT NOT NULL,
        flight_number TEXT NOT NULL,
        flight_date TEXT NOT NULL,
        aircraft_type TEXT NOT NULL,
        seat1 TEXT NOT NULL,
        seat2 TEXT NOT NULL,
        seat3 TEXT NOT NULL,
        created_at TEXT NOT NULL,
        UNIQUE(flight_number, flight_date)
    );`
	_, err := r.db.Exec(query)
	return err
}

// Exists checks if a voucher for a given flight and date already exists.
func (r *sqliteRepository) Exists(flightNumber, flightDate string) (bool, error) {
	var exists bool

	query := `
		SELECT EXISTS(
			SELECT 1 FROM vouchers 
			WHERE flight_number = ? AND flight_date = ?
		)`
	err := r.db.QueryRow(query, flightNumber, flightDate).Scan(&exists)
	if err != nil {
		log.Printf("Error checking existence for flight %s on %s: %v", flightNumber, flightDate, err)
		return false, err
	}

	return exists, nil
}

// Create saves a new voucher assignment to the database.
func (r *sqliteRepository) Create(voucher domain.Voucher) error {
	query := `
        INSERT INTO vouchers (crew_name, crew_id, flight_number, flight_date, aircraft_type, seat1, seat2, seat3, created_at)
        VALUES (:crew_name, :crew_id, :flight_number, :flight_date, :aircraft_type, :seat1, :seat2, :seat3, :created_at);
    `

	dbVoucher := domain.VoucherDB{
		CrewName:     voucher.CrewName,
		CrewID:       voucher.CrewID,
		FlightNumber: voucher.FlightNumber,
		FlightDate:   voucher.FlightDate,
		AircraftType: voucher.AircraftType,
		Seat1:        voucher.Seats[0],
		Seat2:        voucher.Seats[1],
		Seat3:        voucher.Seats[2],
		CreatedAt:    time.Now().UTC().Format(time.RFC3339),
	}

	_, err := r.db.NamedExec(query, &dbVoucher)
	if err != nil {
		log.Printf("Error creating voucher for flight %s: %v", voucher.FlightNumber, err)
	}

	return err
}
