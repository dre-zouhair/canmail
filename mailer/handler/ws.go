package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dre-zouhair/mailer/internal/model"
	service2 "github.com/dre-zouhair/mailer/service"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	_, p, err := conn.ReadMessage()
	if err != nil {
		fmt.Println(err)
		return
	}
	var body Body
	err = json.Unmarshal(p, &body)

	if err != nil {
		err := conn.WriteMessage(websocket.TextMessage, []byte("{'status' : 400}"))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	template, targets := service2.RetrieveBulkData(body.TemplateName)

	if template == nil || targets == nil {
		fmt.Println(errors.New("Unable to retrieve " + body.TemplateName + " data"))
		err := conn.WriteMessage(websocket.TextMessage, []byte("{'status' : 500}"))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	mailServer, err := service2.GetMailServer()
	if err != nil {
		return
	}
	success := make(chan model.Target)
	fails := make(chan model.Target)
	mailServer.ConcurrentBulk(targets, template, success, fails)
	for {
		select {
		case target := <-success:
			err := conn.WriteMessage(websocket.TextMessage, []byte("{\"status\": 200,\"email\":\""+target.Email+"\"}"))
			if err != nil {
				fmt.Println(err)
				return
			}
		case target := <-fails:
			err := conn.WriteMessage(websocket.TextMessage, []byte("{\"status\": 500,\"email\":\""+target.Email+"\"}"))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
