package transport

import (
	"encoding/json"
	"net/http"

	"github.com/shop/internal/schema"
)

// @Summary Получить список пользователей
// @Description Получить список всех пользователей в системе
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {object} schema.User
// @Param   input   body   schema.User 	true "registration"
// @Router /registration [post]
func Registration(w http.ResponseWriter, r *http.Request) {
	var user schema.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
