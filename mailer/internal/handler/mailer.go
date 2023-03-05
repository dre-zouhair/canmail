package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dre-zouhair/mailer/internal/service"
	"net/http"
)

type Body struct {
	TemplateName string `json:"template"`
}

func Bulk(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var body Body
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	template, targets := service.RetrieveBulkData(body.TemplateName)
	if template == nil || targets == nil {
		fmt.Println(errors.New("Unable to retrieve " + body.TemplateName + " data"))
		http.Error(w, "Unable to bulk mail for "+body.TemplateName, http.StatusInternalServerError)
		return
	}
	errs := service.BulkMails(template, targets)

	if len(errs) != 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		rBody := stringifyErrors(errs)
		err := json.NewEncoder(w).Encode(rBody)
		if err != nil {
			obj, _ := json.Marshal(rBody)
			fmt.Println("unable to write stringed errors to the response" + string(obj))
			return
		}
	}

	w.WriteHeader(http.StatusOK)

}

func stringifyErrors(errs []error) []string {
	out := make([]string, 0)
	for _, err := range errs {
		out = append(out, err.Error())
	}
	return out
}
