package controllers

import (
	"log"

	"github.com/gofiber/websocket/v2"
)

var clients = make(map[*websocket.Conn]bool) // Connected clients

func HandleWebSocket(c *websocket.Conn) {
	var (
		mt  int
		msg []byte
		err error
	)
	log.Println(clients)
	for {
		if mt, msg, err = c.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", msg)
		log.Println("Hi")

		if err = c.WriteMessage(mt, []byte("First Message")); err != nil {
			log.Println("write:", err)
			break
		}
	}
}
