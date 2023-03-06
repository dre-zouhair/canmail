package handler

import (
	"encoding/json"
	"github.com/dre-zouhair/mailer/internal/model"
	"github.com/dre-zouhair/mailer/service"
	"net/http"
)

type template struct {
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func SaveTemplate(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var body template
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	templateService := service.NewTemplateService()

	res := templateService.AddTemplate(model.Template{
		Name:    body.Name,
		Body:    body.Body,
		Subject: body.Subject,
	})

	if res != 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Unable to save template")
		return
	}

	w.WriteHeader(http.StatusOK)

}
