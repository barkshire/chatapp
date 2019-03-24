package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpointHandler(w http.ResponseWriter, r *http.Request) {

	// workaround to avoid CORS error
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade to a websocket connection
	// returns a pointer to a websocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Client Connected.")

	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		fmt.Println(err)
	}

	reader(ws)
}

// listens to incoming msg to websocket endpoint
func reader(conn *websocket.Conn) {
	for {

		// read in msg to buffer
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// print read msg
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func setupRoutes() {
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/ws", wsEndpointHandler)
}

// websocket upgrader to hold Read/Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
