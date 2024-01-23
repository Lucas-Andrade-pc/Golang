package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Response interface {
	getResponse() string
}

func (r Results) getResponse() string {
	fmt.Printf("Status body: %d\n", r.Status)
	return fmt.Sprintln(r.Body)
}

type Results struct {
	Status int
	Body   string `json:"error"` // armazena os dados de erro vindo do json analizado example: {error: "Full authentication is required to access this resource"}
}

type Page struct {
	Body map[string]int `json:"qnt"`
}

func main() {
	args := os.Args // receber argumentos passado por linhda de comando go run main <arg1>

	if len(args) < 2 { // verficia o tamanho do array de argumentos passado
		fmt.Println("Usage: ./main <url>")
		os.Exit(1)
	}
	res, err := doRequest(args[1])
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	if res == nil {
		fmt.Printf("No response\n")
	}

	fmt.Printf("Response: %s", res.getResponse())
}
func doRequest(requestUrl string) (Response, error) {

	_, err := url.ParseRequestURI(requestUrl) // faz uma analisa de uma url HTTP|HTTPS

	if err != nil {
		fmt.Printf("URL is in invalid format %s\n", requestUrl)
		os.Exit(1)
	}

	response, err := http.Get(requestUrl) // faz uma request a url passada

	if err != nil {
		log.Fatal(err)

	}

	defer response.Body.Close() // defer adia a execução de uma função até que a função circundante retorne.

	body, err := io.ReadAll(response.Body) // le todo o retorno do body retornado
	if err != nil {
		log.Fatal(err)
	}
	// switch response.StatusCode {
	// case 200:
	// 	fmt.Printf("HTTP Status Code: %d\nBody: %ss", response.StatusCode, body)
	// default:
	// 	fmt.Printf("Error")
	// 	os.Exit(1)
	// }
	var result Results // declarando variavel do tipo Result
	if response.StatusCode != 200 {
		result.Status = response.StatusCode
		err = json.Unmarshal(body, &result) // analisa os dados em JSON e armazena o resultado no valor apontado, ou seja, passei os dados body para a função e ele armazenou em nosso estrutura
		if err != nil {
			log.Fatal(err)
		}
		return result, nil
	} else {
		return result, nil
	}
}
