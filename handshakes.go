package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/json"
	"errors"
	"net"
)

/*
	A brief description on how the protocol works:
	1-	generate a 4096bit RSA key
	2-	send a JSON version of it to the peer
	3-	check if the peer has responded with "KEY OK"
		3.5-	if it has not, then stop the handshake
	4-	generate a random sequence of characters with
		length of ten
	5-	send the sequence to the peer
	6-	receive a buffer from the client
	7-	try to decipher the buffer using your private
		key
		7.5-	if you cannot, then stop the handshake
	8-	check if the decipher matches the sequence that
		was generated at step 4
		8.5-	if not, then stop the handshake
*/
func handshake(c net.Conn, priv *rsa.PrivateKey) error {
	serialized_pub := serializePubKey(&priv.PublicKey)
	response, err := sendAndRead(serialized_pub, c)

	if err != nil {
		return err
	}

	if response != "KEY OK" {
		return errors.New("bad response to key")
	}

	test_string := generateRandomString(10)
	response, err = sendAndRead(test_string, c)

	if err != nil {
		return err
	}

	decipher, err := rsa.DecryptOAEP(
		sha512.New(),
		rand.Reader,
		priv,
		[]byte(response),
		[]byte("RESPONSE"),
	)

	if err != nil {
		return err
	}

	if string(decipher) != test_string {
		return errors.New("the decipher does not match original string")
	}

	return nil
}

func handleHandshake(c net.Conn) (*rsa.PublicKey, error) {
	json_pub, err := readAndSend("KEY OK", c)

	if err != nil {
		return nil, err
	}

	pub := &rsa.PublicKey{}
	err = json.Unmarshal([]byte(json_pub), pub)

	if err != nil {
		return nil, err
	}

	test_string, err := read(c)

	if err != nil {
		return nil, err
	}

	cipher, err := rsa.EncryptOAEP(
		sha512.New(),
		rand.Reader,
		pub,
		[]byte(test_string),
		[]byte("RESPONSE"),
	)

	if err != nil {
		return nil, err
	}

	send(string(cipher), c)

	return pub, nil
}
