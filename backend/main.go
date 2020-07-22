package main

import (
	"fmt"
	"net/http"

	"github.com/robatussum/ccb/backend/pkg/ws"
)

func serveWs(pool *ws.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("ws endpoint hit!")
	conn, err := ws.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}

	client := &ws.Client{
		Conn: conn,
		Pool: ws.Pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := ws.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Distributed chat app v0.01")
	setupRoutes()

	http.ListenAndServe(":9000", nil)
}
