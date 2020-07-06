package controllers

import (
	"net/http"

	"github.com/mlph-kvillegas/events-reservation-system-backend/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To Events Reservation System")
}
