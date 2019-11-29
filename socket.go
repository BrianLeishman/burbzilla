package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

func socket(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	check(err)
	defer c.Close()

	for {
		j, err := json.Marshal(sensors)
		check(err)

		c.WriteMessage(websocket.TextMessage, j)

		time.Sleep(time.Millisecond * 75)
	}
}
