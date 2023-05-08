package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/aronkst/grpc-vs-rest-go/data"
)

func API(writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}

	data := data.Data{}
	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("data (REST Server): %v\n", data)

	dataBody, err := json.Marshal(&data)
	if err != nil {
		log.Fatal(err)
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(dataBody)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", API)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
