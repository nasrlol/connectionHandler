package main

import (
	"net/http"
	"time"
	"log"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		//		return r.Host == "nsrddyn.com"	
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request){
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading to WebSocket :%v\n", err)
		return
	}
	defer conn.Close()

	for {
		dataType, data, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading WebSocket message: %v\n", err)
			break
		}

		log.Printf("Received: %s\n", data)

		if err := conn.WriteMessage(dataType, data); err != nil {
			log.Printf("Error writing WebSocket message: %v\n", err)
			break
		}
	}
}

func main(){

	server := &http.Server{
		Addr: ":8085",
		Handler: http.HandlerFunc(wsHandler), 
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil{
		log.Fatalf("Failed to listen and serve the server %v\n", err)
	}
}

