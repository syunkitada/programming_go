package main

import (
	"bufio"
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:9090", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer func() {
			os.Exit(0)
			close(done)
		}()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
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
		scanner.Scan()
		err := c.WriteMessage(websocket.TextMessage, []byte(scanner.Text()))
		if err != nil {
			log.Println("write:", err)
			return
		}
	}
}
