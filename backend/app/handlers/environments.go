package handlers

import (
	"encoding/json"
	"go-app/models"
	"go-app/services"
	"go-app/utils"
	"net/http"
)

func CreateEnvironment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var env models.Environment
	if err := json.NewDecoder(r.Body).Decode(&env); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate input
	if errors := utils.ValidateEnvironment(env); len(errors) > 0 {
		http.Error(w, errors, http.StatusBadRequest)
		return
	}

	// Deploy environment using Helm
	if err := services.DeployEnvironment(env); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Environment deployed successfully"})
}
