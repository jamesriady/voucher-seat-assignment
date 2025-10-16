package handler

import (
	"encoding/json"
	"net/http"

	"backend/internal/domain"
	internalError "backend/internal/error"
	"backend/internal/service"

	"github.com/go-playground/validator/v10"
)

// VoucherHandler handles the HTTP requests for vouchers.
type VoucherHandler struct {
	service  service.VoucherService
	validate *validator.Validate
}

// NewVoucherHandler creates a new VoucherHandler.
func NewVoucherHandler(
	s service.VoucherService,
	validate *validator.Validate,
) *VoucherHandler {
	return &VoucherHandler{service: s, validate: validate}
}

func (h *VoucherHandler) CheckVoucher(w http.ResponseWriter, r *http.Request) {
	var req domain.CheckVoucherRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := ValidateStruct(h.validate, req); err != nil {
		respondJSON(w, internalError.GetStatus(err), err)
		return
	}

	exists, sErr := h.service.CheckVoucherExists(req.FlightNumber, req.Date)
	if sErr != nil {
		statusCode := internalError.GetStatus(sErr)
		respondError(w, statusCode, sErr.Error())
		return
	}

	respondJSON(w, http.StatusOK, map[string]bool{"exists": exists})
}

func (h *VoucherHandler) GenerateVoucher(w http.ResponseWriter, r *http.Request) {
	var req domain.GenerateVouchersRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := ValidateStruct(h.validate, req); err != nil {
		respondJSON(w, internalError.GetStatus(err), err)
		return
	}

	res, sErr := h.service.GenerateVouchers(req)
	if sErr != nil {
		statusCode := internalError.GetStatus(sErr)
		respondError(w, statusCode, sErr.Error())
		return
	}

	respondJSON(w, http.StatusCreated, res)
}
