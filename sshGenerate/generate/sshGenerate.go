package ssh

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/pem"
	"fmt"
)

func GenerateKeys() ([]byte, []byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Read, 4096)
	if err != nil {
		return nil, nil, fmt.Errorf("erro create private key -> %s", err)
	}

	privatrKeyPem := &pem.Block

	return nil, nil, nil
}
