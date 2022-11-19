package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func main() {
	go subToITH()

	for {

	}
}

func subToITH() {
	// Connect to a server
	nc, _ := nats.Connect("nats://95.165.107.100:4222")
	defer nc.Drain()

	// // Simple Publisher
	// nc.Publish("oleja228 bot", []byte("Hello World"))

	// Responding to a request message
	nc.Subscribe("ith", func(m *nats.Msg) {
		m.Respond([]byte("Безусов принял"))
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	// Close connection
	// nc.Close()

}
