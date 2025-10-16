package service

import (
	"errors"
	"fmt"
	"log"
	"math/rand"

	"backend/internal/domain"
	internalError "backend/internal/error"
	"backend/internal/repository"
)

// VoucherService defines the business logic for voucher operations.
type VoucherService interface {
	CheckVoucherExists(flightNumber, flightDate string) (bool, error)
	GenerateVouchers(request domain.GenerateVouchersRequest) (domain.GenerateVouchersResponse, error)
}

type voucherService struct {
	repo repository.VoucherRepository
}

// NewVoucherService creates a new voucher service.
func NewVoucherService(r repository.VoucherRepository) VoucherService {
	return &voucherService{repo: r}
}

// CheckVoucherExists checks if a voucher already exists for a flight.
func (s *voucherService) CheckVoucherExists(flightNumber, flightDate string) (bool, error) {
	exists, err := s.repo.Exists(flightNumber, flightDate)
	if err != nil {
		log.Printf("failed to check for existing vouchers: %v \n", err)
		return true, errors.New("failed to check for existing vouchers")
	}

	return exists, nil
}

// GenerateVouchers creates unique seat vouchers and saves them.
func (s *voucherService) GenerateVouchers(request domain.GenerateVouchersRequest) (domain.GenerateVouchersResponse, error) {
	exists, err := s.repo.Exists(request.FlightNumber, request.FlightDate)
	if err != nil {
		log.Printf("failed to check for existing vouchers: %v \n", err)
		return domain.GenerateVouchersResponse{}, errors.New("failed to check for existing vouchers")
	}
	if exists {
		return domain.GenerateVouchersResponse{
			Success: false,
			Error:   "Vouchers already exist for this flight",
		}, nil
	}

	// Generate the seats based on aircraft type.
	seats, err := generateRandomSeats(request.AircraftType)
	if err != nil {
		return domain.GenerateVouchersResponse{}, err
	}

	voucher := domain.Voucher{
		CrewName:     request.CrewName,
		CrewID:       request.CrewID,
		FlightNumber: request.FlightNumber,
		FlightDate:   request.FlightDate,
		AircraftType: request.AircraftType,
		Seats:        seats,
	}

	if err := s.repo.Create(voucher); err != nil {
		log.Printf("failed to create voucher: %v \n", err)
		return domain.GenerateVouchersResponse{}, errors.New("failed to save the generated vouchers")
	}

	return domain.GenerateVouchersResponse{
		Success: true,
		Seats:   seats,
	}, nil
}

// generateRandomSeats creates unique seats based on the aircraft layout.
func generateRandomSeats(aircraftType domain.AircraftType) ([]string, error) {
	var (
		rows        int
		seatsPerRow []string
	)

	switch aircraftType {
	case domain.ATR:
		rows = domain.ATRMaxRows
		seatsPerRow = domain.ATRSeatsPerRow
	case domain.Airbus320:
		rows = domain.Airbus320MaxRows
		seatsPerRow = domain.Airbus320SeatsPerRow
	case domain.Boeing737Max:
		rows = domain.Boeing737MaxMaxRows
		seatsPerRow = domain.Boeing737MaxSeatsPerRow
	default:
		return nil, internalError.ErrInvalidAircraft
	}

	// Create a pool of all possible seats.
	var seatPool []string
	for r := 1; r <= rows; r++ {
		for _, s := range seatsPerRow {
			seatPool = append(seatPool, fmt.Sprintf("%d%s", r, s))
		}
	}

	// Shuffle the pool
	rand.Shuffle(len(seatPool), func(i, j int) {
		seatPool[i], seatPool[j] = seatPool[j], seatPool[i]
	})

	if len(seatPool) < domain.MaxVoucherPerFlight {
		return nil, errors.New("not enough seats available to generate vouchers")
	}

	return seatPool[:domain.MaxVoucherPerFlight], nil
}
