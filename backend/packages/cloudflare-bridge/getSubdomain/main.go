package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const API_URL = "https://api.cloudflare.com/client/v4/zones/%zone_id%/dns_records?type=A"

// Serverless struct

type Request struct {
}

type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       ResponseBody      `json:"body,omitempty"`
}

// Cloudflare API response struct

type ResponseBody struct {
	Result []ResultResponse `json:"result"`
}

type ApiResponse struct {
	Result  []ResultResponse `json:"result"`
	Success bool             `json:"success"`
}

type ResultResponse struct {
	Name      string `json:"name"`
	Content   string `json:"content"`
	CreatedOn string `json:"created_on"`
}

func Main(in Request) (*Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest(
		"GET",
		strings.Replace(API_URL, "%zone_id%", os.Getenv("CLOUDFLARE_ZONE_ID"), 1),
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("X-Auth-Key", os.Getenv("CLOUDFLARE_AUTH_KEY"))
	req.Header.Add("X-Auth-Email", os.Getenv("CLOUDFLARE_AUTH_EMAIL"))

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var jsonResp ApiResponse
	json.Unmarshal([]byte(data), &jsonResp)

	return &Response{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: ResponseBody{jsonResp.Result},
	}, nil
}
