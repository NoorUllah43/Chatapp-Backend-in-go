package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	From    int    `json:"from"`
	To      int    `json:"to"`
	Message string `json:"message"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients map[string]*websocket.Conn

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}
	defer conn.Close()
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		var msg Message
		err = json.Unmarshal(data, &msg)
		if err != nil {
			fmt.Println("JSON unmarshal error:", err)
		}

		fmt.Println(msg)
		
		err = conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			fmt.Println("Error writing message:", err)
			break
		}

	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	fmt.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
