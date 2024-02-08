package server

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

func StartServer(privatePem []byte, authorizedKeys []byte) error {
	authorizedKeysMap := map[string]bool{}
	for len(authorizedKeys) > 0 {
		pubKey, _, _, rest, err := ssh.ParseAuthorizedKey(authorizedKeys) //ParseAuthorizedKey analisa uma chave pública de um arquivoauthorized_keys
		if err != nil {
			return fmt.Errorf("error public key -> %s", err)
		}

		authorizedKeysMap[string(pubKey.Marshal())] = true //Marshal retorna os dados da chave serializada no formato de ligação SSH, com o prefixo do nome.
		authorizedKeys = rest
	}
	config := &ssh.ServerConfig{
		PublicKeyCallback: func(c ssh.ConnMetadata, pubKey ssh.PublicKey) (*ssh.Permissions, error) {
			if authorizedKeysMap[string(pubKey.Marshal())] {
				return &ssh.Permissions{
					// Record the public key used for authentication.
					Extensions: map[string]string{
						"pubkey-fp": ssh.FingerprintSHA256(pubKey),
					},
				}, nil
			}
			return nil, fmt.Errorf("unknown public key for %q", c.User())
		},
	}
	// privateBytes, err := os.ReadFile("id_rsa")
	// if err != nil {
	// 	log.Fatal("Failed to load private key: ", err)
	// }

	private, err := ssh.ParsePrivateKey(privatePem)
	if err != nil {
		log.Fatal("Failed to parse private key: ", err)
	}
	config.AddHostKey(private)

	// Once a ServerConfig has been configured, connections can be
	// accepted.
	listener, err := net.Listen("tcp", "0.0.0.0:2022")
	if err != nil {
		log.Fatal("failed to listen for connection: ", err)
	}
	nConn, err := listener.Accept()
	if err != nil {
		log.Fatal("failed to accept incoming connection: ", err)
	}
	return nil
}
