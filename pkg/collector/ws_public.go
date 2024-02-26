package collector

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func ConnectWS(ctx context.Context, topic string, address string, chanelData chan map[string]interface{}) {
	c, _, err := websocket.DefaultDialer.Dial(address, nil)
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	defer c.Close()

	fmt.Println("Connected.")

	subscribe(c, topic)

	for {

		select {
		case <-ctx.Done():
			return
		default:
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("Failed to read message:", err)
				continue
			}
			data, err := receive(string(message))
			chanelData <- data
		}
	}
}

func subscribe(c *websocket.Conn, topic string) {
	subscribeRequest := map[string]interface{}{
		"op":   "subscribe",
		"args": []string{topic},
	}
	message, err := json.Marshal(subscribeRequest)
	if err != nil {
		log.Println("Failed to marshal JSON:", err)
		return
	}
	err = c.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("Failed to send subscribe message:", err)
		return
	}
}

func keepAlive(c *websocket.Conn) {
	for {
		err := c.WriteMessage(websocket.PingMessage, nil)
		if err != nil {
			log.Println("Ping failed:", err)
			return
		}
		time.Sleep(time.Second * 10)
	}
}
