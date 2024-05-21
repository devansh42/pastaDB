package main

import (
	"log"
	"net"
	"strconv"
)

// Implements PastaDB server

type server struct {
	port int
}

func (s server) start() {
	addr := strconv.Itoa(s.port)
	addr = "localhost:" + addr
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("error while starting tcp server: " + err.Error())
	}
	log.Print("Server is ready to accept connections at: ", addr)
	for {
		incomingConn, err := listener.Accept()
		if err != nil {
			log.Print("error while accepting new connection: ", err.Error())
		}

		// handling new connection in separate
		// go-routine
		go handleNewConnection(incomingConn)
	}
}
