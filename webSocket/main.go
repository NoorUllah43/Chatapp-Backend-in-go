package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"websocket/mongodb"

	"github.com/gorilla/websocket"
)

type Message struct {
	SenderID   int    `json:"senderID"`
	ReceiverID int    `json:"receiverID"`
	Message    string `json:"message"`
	Time       string `json:"time"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[int]*websocket.Conn)
var mutex = &sync.Mutex{}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading:", err)
		return
	}
	query := r.URL.Query()
	userID := query.Get("userId")

	id, err := strconv.Atoi(userID)
	if err != nil {
		fmt.Println("Invalid userId:", err)
		return
	}
	mutex.Lock()
	clients[id] = conn
	mutex.Unlock()

	defer conn.Close()
	messageCollection := mongodb.GetCollection("Messages")
	
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

		var receiverConn = clients[msg.ReceiverID]

		err = receiverConn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			fmt.Println("Error writing message:", err)
			break
		}

		_, err = messageCollection.InsertOne(context.Background(), msg)
		if err != nil {
			fmt.Println("Error saving message to MongoDB:", err)
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
