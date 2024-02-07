package main

import (
	"fmt"
	"os"
	ssh "sshgenerate/generate"
)

func main() {
	var (
		privatePem []byte
		publicKey  []byte
		err        error
	)

	var a *int
	*a = 10
	fmt.Println(*a)
	if privatePem, publicKey, err = ssh.GenerateKeys(); err != nil {
		fmt.Printf("Error -> %s", err)
		os.Exit(1)
	}
	if err = os.WriteFile("my.pem", privatePem, 0600); err != nil {
		fmt.Printf("Erro -> %s", err)
		os.Exit(1)
	}
	if err = os.WriteFile("my.pub", publicKey, 0644); err != nil {
		fmt.Printf("Erro -> %s", err)
		os.Exit(1)
	}

}
