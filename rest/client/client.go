package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	jsonPerson, err := os.Open("./rest/client/persons.json")

	if err != nil {
		log.Fatalln("error opening file", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:4000", jsonPerson)

	if err != nil {
		log.Fatalln("request error", err)
	}

	defer req.Body.Close()

	res, err := http.DefaultClient.Do(req)

	b, _ := io.ReadAll(res.Body)
	fmt.Printf("status=%d body=%s\n", res.StatusCode, string(b))
}
