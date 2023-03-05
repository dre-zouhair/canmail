package main

import (
	"fmt"
	"github.com/dre-zouhair/mailer/internal/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte("mailer is UP"))
		if err != nil {
			return
		}
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/bulk", handler.Bulk)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("unable to start the server")
		return
	}
}
