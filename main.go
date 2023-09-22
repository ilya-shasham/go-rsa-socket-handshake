package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func server(wg *sync.WaitGroup) {
	listener, err := net.Listen("tcp", "localhost:9090")
	handleError(err)
	client, err := listener.Accept()
	handleError(err)
	priv, err := rsa.GenerateKey(rand.Reader, 4096)
	handleError(err)
	err = handshake(client, priv)
	handleError(err)
	client_pub, err := handleHandshake(client)
	handleError(err)
	fmt.Println("Got client's public key: ", client_pub)
	wg.Done()
}

func client(wg *sync.WaitGroup) {
	connection, err := net.Dial("tcp", "localhost:9090")
	handleError(err)
	server_pub, err := handleHandshake(connection)
	handleError(err)
	fmt.Println("Got server's public key: ", server_pub)
	priv, err := rsa.GenerateKey(rand.Reader, 4096)
	handleError(err)
	err = handshake(connection, priv)
	handleError(err)
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go server(wg)

	// Just there to make sure
	// the server has started
	// working
	time.Sleep(time.Second)

	wg.Add(1)
	go client(wg)

	wg.Wait()
}
