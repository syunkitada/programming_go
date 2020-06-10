package main

import (
	"bufio"
	"flag"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:9090", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	log.Printf("connected")

	done := make(chan struct{})

	go func() {
		defer func() {
			os.Exit(0)
			close(done)
		}()
		for {
			log.Println("Waiting read message")
			_, message, err := c.ReadMessage()
			if err != nil {
				// Server側Close時のエラー: read: websocket: close 1006 (abnormal closure): unexpected EOF
				log.Println("read:", err)
				return
			}
			log.Println(string(message))
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		log.Println("Waiting write message")
		scanner.Scan()
		err := c.WriteMessage(websocket.TextMessage, []byte(scanner.Text()))
		if err != nil {
			log.Println("write:", err)
			return
		}
	}
}
