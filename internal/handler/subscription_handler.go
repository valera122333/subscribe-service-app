package handler

import (
	"encoding/json"
	"net/http"
	"subscriptions-service/internal/model"
	"subscriptions-service/internal/service"

	"github.com/google/uuid"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var sub model.Subscription
	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	id, err := h.service.Create(r.Context(), sub)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id.String()})
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	serviceName := r.URL.Query().Get("service_name")

	var userID uuid.UUID
	var err error
	if userIDStr != "" {
		userID, err = uuid.Parse(userIDStr)
		if err != nil {
			http.Error(w, "invalid user_id", 400)
			return
		}
	}

	list, err := h.service.ListFiltered(r.Context(), userID, serviceName)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")

	id, _ := uuid.Parse(idStr)

	sub, err := h.service.Get(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	json.NewEncoder(w).Encode(sub)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")

	id, _ := uuid.Parse(idStr)

	err := h.service.Delete(r.Context(), id)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
func (h *Handler) Total(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	serviceName := r.URL.Query().Get("service_name")
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	var userID uuid.UUID
	var err error
	if userIDStr != "" {
		userID, err = uuid.Parse(userIDStr)
		if err != nil {
			http.Error(w, "invalid user_id", 400)
			return
		}
	}

	total, err := h.service.Total(r.Context(), userID, serviceName, from, to)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{
		"total_price": total,
	})
}
