package main

import (
	"net"
	"time"
)

// Reads a maximum of 10kbs from the given
// connection as a buffer, and then returns
// the buffer as a string.
// the function returns an error value relating
// to socket read timeout if the wait time
// exceeds five seconds.
func read(c net.Conn) (string, error) {
	err := c.SetReadDeadline(time.Now().Add(5 * time.Second))

	if err != nil {
		return "", err
	}

	data_buffer := make([]byte, 1024*10)
	n_read, err := c.Read(data_buffer)

	return string(data_buffer[:n_read]), err
}

// Sends data to a connection.
func send(message string, c net.Conn) {
	c.Write([]byte(message))
}

// First sends the given message, then
// waits for and reads a response.
func sendAndRead(message string, c net.Conn) (string, error) {
	send(message, c)
	return read(c)
}

// First reads from the connection, then
// sends a response.
func readAndSend(message string, c net.Conn) (string, error) {
	result, err := read(c)
	send(message, c)
	return result, err
}
