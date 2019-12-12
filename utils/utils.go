package utils

import (
	"log"
)

func ListenWebsocket(ws WebsocketActor, ch chan string) {
	for {
		_, message, err := ws.Connection().ReadMessage()
		if err != nil {
			log.Printf("Err - %s said: %s\n", ws.Name(), err.Error())
			return
		}
		log.Printf("%s listened: %s", ws.Name(), message)
		// send message
		ch <- string(message)
	}
}
func SendWebsocketMessage(ws WebsocketActor, ch chan string) {
	for {
		message := <-ch
		log.Printf("%s said: %s", ws.Name(), message)
		if err := ws.Connection().WriteMessage(1, []byte(message)); err != nil {
			log.Println(err)
			return
		}
	}
}
