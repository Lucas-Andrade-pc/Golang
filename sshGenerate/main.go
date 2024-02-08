package main

import (
	"fmt"
	"os"
	ssh "sshgenerate/generate"
	"sshgenerate/server"
)

func main() {
	var (
		privatePem  []byte
		publicKey   []byte
		filePrivate []byte
		filePublic  []byte
		err         error
	)

	if privatePem, publicKey, err = ssh.GenerateKeys(); err != nil {
		fmt.Printf("Error -> %s", err)
		os.Exit(1)
	}
	if err = os.WriteFile("mykey.pem", privatePem, 0600); err != nil {
		fmt.Printf("Erro -> %s", err)
		os.Exit(1)
	}
	if err = os.WriteFile("mykey.pub", publicKey, 0644); err != nil {
		fmt.Printf("Erro -> %s", err)
		os.Exit(1)
	}

	if filePrivate, err = os.ReadFile("mykey.pem"); err != nil {
		fmt.Printf("ReadFile error -> %s", err)
		os.Exit(1)
	}
	if filePublic, err = os.ReadFile("mykey.pub"); err != nil {
		fmt.Printf("ReadFile error -> %s", err)
		os.Exit(1)
	}

	if err := server.StartServer(filePrivate, filePublic); err != nil {
		fmt.Printf("Error start server -> %s", err)
		os.Exit(1)
	}

}
