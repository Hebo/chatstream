package main

import (
	"fmt"
	"log"
	"net"

	"time"
)

func main() {
	fmt.Println("hello")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello world")
	// })

	ln, err := net.Listen("tcp4", ":9999")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go sendMessages(conn)
	}

}

func sendMessages(conn net.Conn) {
	count := 0
	for {
		fmt.Fprintf(conn, "hello world\n")

		count += 1
		if count%10 == 0 {
			fmt.Printf("%d messages sent\n", count)
		}

		time.Sleep(1 * time.Second)
	}
}
