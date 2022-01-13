package orders

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type DataStruct struct {
	Data Data `json:"data"`
}

type Data struct {
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	Amount int `json:"amount"`
}

func PaymongoService(ctx *gin.Context) {
	// Request url and response body
	URL := "https://api.paymongo.com/v1/sources"
	attrs := DataStruct{}
	attrs.Data.Attributes.Amount = 10000
	fmt.Println(attrs)
	fmt.Println("hereherehere2424242")
	postBody, _ := json.Marshal(attrs)
	os.Stdout.Write(postBody)
	fmt.Println("hereherehere123123")
	responseBody := bytes.NewBuffer(postBody)

	// // Create a Basic string by appending string access token
	BasicAuth := "Basic " + "pk_test_hAaLoSxFQ7dF8771yv4owcAU"

	// // Create a new request using http
	req, err := http.NewRequest("Post", URL, responseBody)

	// // add authorization header to the req
	req.Header.Add("Authorization", BasicAuth)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// // Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println(string([]byte(body)))
}
