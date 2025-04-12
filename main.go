package main 

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
	"log"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
	EnableCompression: true,

}

func wsHandler(w http.ResponseWriter, r *http.Request){
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading", err)
		return
	}
	defer conn.Close()
	for {
		sendData := Sys.CPU()

	}
}

func main(){
	server := &http.Server{
		Addr: ":8085",
		Handler: wsHandler,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	http.HandleFunc("/ws", wsHandler)
	fmt.Println("Websocket server started on :8080")
	log.Fatal(server.ListenAndServe)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
