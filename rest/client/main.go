package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/aronkst/grpc-vs-rest-go/data"
)

func main() {
	countString := flag.String("count", "", "")
	flag.Parse()

	countInt, err := strconv.Atoi(*countString)
	if err != nil {
		log.Fatal(err)
	}

	count := int32(countInt)

	for i := int32(1); i <= count; i++ {
		dataInput := data.Data{
			Value1: i,
			Value2: i,
			Value3: i,
			Value4: i,
			Value5: i,
		}

		dataBody, err := json.Marshal(dataInput)
		if err != nil {
			log.Fatal(err)
		}

		dataReader := bytes.NewBuffer(dataBody)

		resp, err := http.Post("http://localhost:8080/", "application/json", dataReader)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		data := data.Data{}
		err = json.Unmarshal(body, &data)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("data (REST Client): %v, i: %d\n", data, i)
	}
}
