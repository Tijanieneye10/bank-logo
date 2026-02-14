package handlers

import (
	"bankplay/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetBanks(w http.ResponseWriter, r *http.Request) {
	banks := models.ListOfBanks
	err := json.NewEncoder(w).Encode(banks)

	if err != nil {
		return
	}
}

func GetSingleBank(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	for _, bank := range models.ListOfBanks {
		if bank.Code == code {
			err := json.NewEncoder(w).Encode(bank)
			if err != nil {
				fmt.Println(err)
				return
			}
			return
		}
	}

	http.NotFound(w, r)
}
