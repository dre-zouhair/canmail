package handler

import (
	"encoding/json"
	"github.com/dre-zouhair/mailer/internal/model"
	"github.com/dre-zouhair/mailer/service"
	"net/http"
)

func SaveTarget(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var body model.Target
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	templateService := service.NewTargetService()
	res := templateService.AddTarget(body)

	if res != 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Unable to save target")
		return
	}

	w.WriteHeader(http.StatusOK)

}
