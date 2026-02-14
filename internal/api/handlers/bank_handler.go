package handlers

import (
	"bankplay/internal/models"
	"encoding/json"
	"net/http"
)

func GetBanks(w http.ResponseWriter, r *http.Request) {
	banks := models.ListOfBanks
	json.NewEncoder(w).Encode(banks)
}
