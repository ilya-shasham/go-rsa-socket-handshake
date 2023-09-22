package main

import (
	"crypto/rsa"
	"encoding/json"
)

func serializePubKey(pub *rsa.PublicKey) string {
	var marshalled, _ = json.Marshal(pub)
	return string(marshalled)
}
