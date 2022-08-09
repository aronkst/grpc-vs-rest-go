package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aronkst/grpc-vs-rest-go/data"
)

func Api(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	data := data.Data{}
	json.Unmarshal(body, &data)

	fmt.Printf("data (REST Server): %v\n", data)

	dataBody, err := json.Marshal(&data)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataBody)
}

func main() {
	http.HandleFunc("/", Api)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
