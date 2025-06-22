package utils

import (
	"encoding/json"
	"github.com/farahsrw/manageProject/models"
	"net/http"
)

func RespondError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(models.ErrorResponse{Message: message})
}
