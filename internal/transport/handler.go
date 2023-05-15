package transport

import (
	"encoding/json"
	"net/http"

	"github.com/shop/internal/schema"
)

// Registration godoc
// @Summary Registration
// @Description Registration
// @Tags registration
// @ID registration
// @Accept  json
// @Produce  json
// @Param registration body schema.User true "Registration Input"
// @Success 200 {object} schema.Error
// @Router /registration [post]
func Registration(w http.ResponseWriter, r *http.Request) {
	var user schema.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// Обработайте ошибку некорректных данных
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	response := struct {
		Message string `json:"message"`
	}{
		Message: "User registered successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
