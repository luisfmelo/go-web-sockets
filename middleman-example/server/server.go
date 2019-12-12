package main

import (
	"log"
	"net/http"
	u "pkg/ws/utils"
)

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	var msgChannel = make(chan string)

	wsServer := u.NewWebsocketServer()
	wsServer.UpgradeConnection(w, r)

	// As a websocket server - start listening to messages from client
	go u.ListenWebsocket(wsServer, msgChannel)
	// As a websocket server - redirect messages from server to client
	go u.SendWebsocketMessage(wsServer, msgChannel)
}

func main() {

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status": "alive", "message": "I'm the server!'"}`))
	})

	http.HandleFunc("/ws", wsEndpoint)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
