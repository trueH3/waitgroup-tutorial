package main

import (
	"log"
	"net"
	"sync/atomic"
	"time"
)

func main() {

	//thats how we create tcp server
	li, err := net.Listen("tcp", ":8008")

	if err != nil {
		log.Fatalf("Could not create server %v", err)
	}

	var connections int32
	for {
		conn, err := li.Accept()

		if err != nil {
			continue
		}
		connections++

		go func ()  {
			//server connection

			defer func ()  {
				conn.Close()
				atomic.AddInt32(&connections, -1)
			}()

			if atomic.LoadInt32(&connections) > 3 {
				return
			}

			time.Sleep(time.Second)
			conn.Write([]byte("success"))
		}()
	}
}