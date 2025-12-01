package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"avito_project/course-go-avito-Vodyiara/internal/model"
	"avito_project/course-go-avito-Vodyiara/internal/service"

	"github.com/go-chi/chi/v5"
)

type CourierHandler struct {
	service service.CourierService
}

func NewCourierHandler(service service.CourierService) *CourierHandler {
	return &CourierHandler{
		service: service,
	}
}

func (h *CourierHandler) CreateCourier(w http.ResponseWriter, r *http.Request) {
	var req model.CreateCourierRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	courier, err := h.service.CreateCourier(r.Context(), &req)
	if err != nil {
		h.handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(courier)
}

func (h *CourierHandler) GetCourier(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	courier, err := h.service.GetCourier(r.Context(), id)
	if err != nil {
		h.handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(courier)
}

func (h *CourierHandler) GetAllCouriers(w http.ResponseWriter, r *http.Request) {
	couriers, err := h.service.GetAllCouriers(r.Context())
	if err != nil {
		h.handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(couriers)
}

func (h *CourierHandler) UpdateCourier(w http.ResponseWriter, r *http.Request) {
	var req model.UpdateCourierRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateCourier(r.Context(), &req); err != nil {
		h.handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("courier updated successfully"))
}

func (h *CourierHandler) handleError(w http.ResponseWriter, err error) {
	log.Printf("Handler error: %v", err)

	switch {
	case errors.Is(err, model.ErrCourierNotFound):
		http.Error(w, err.Error(), http.StatusNotFound)
	case errors.Is(err, model.ErrInvalidName),
		errors.Is(err, model.ErrInvalidPhone),
		errors.Is(err, model.ErrInvalidStatus),
		errors.Is(err, model.ErrIDRequired),
		errors.Is(err, model.ErrInvalidID),
		errors.Is(err, model.ErrPhoneAlreadyExists):
		http.Error(w, err.Error(), http.StatusBadRequest)
	default:
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
