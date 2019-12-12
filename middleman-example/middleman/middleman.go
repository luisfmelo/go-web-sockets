package main

import (
	"log"
	"net/http"
	u "pkg/ws/utils"
)

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	var messageToClient = make(chan string)
	var messageToServer = make(chan string)

	// ----- Start Websocket server connection to client -----
	wsServer := u.NewWebsocketServer()
	wsServer.UpgradeConnection(w, r)
	// ----- Start Websocket client connection to server -----
	wsClient := u.NewWebsocketClient("localhost:8080")

	//--------
	// As a websocket server - start listening to messages from client
	go u.ListenWebsocket(wsServer, messageToServer)
	// As a websocket client - start listening to messages from server
	go u.ListenWebsocket(wsClient, messageToClient)

	// As a websocket server - redirect messages from server to client
	go u.SendWebsocketMessage(wsServer, messageToClient)
	// As a websocket client - redirect messages from client to server
	go u.SendWebsocketMessage(wsClient, messageToServer)
}

func main() {

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status": "alive", "message": "I'm just the middleman'"}`))
	})

	http.HandleFunc("/ws", wsEndpoint)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
