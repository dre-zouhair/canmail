package main

import (
	"fmt"
	"github.com/dre-zouhair/mailer/handler"
	"net"
	"net/http"
	"time"
)

func main() {
	Up()
	defer Down()

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
	http.HandleFunc("/template/save", handler.SaveTemplate)
	http.HandleFunc("/template", handler.GetTemplates)
	http.HandleFunc("/target/save", handler.SaveTarget)
	http.HandleFunc("/ws", handler.WsHandler)

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", "", "8080"))
	if err != nil {
		fmt.Printf("Error creating listener: %v\n", err)
		return
	}

	server := http.Server{
		Handler:      nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 0 * time.Second,
	}

	if err := server.Serve(listener); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
