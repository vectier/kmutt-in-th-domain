package main

import (
	"fmt"
)

type Response struct {
	StatusCode int `json:"status_code"`
}

func Main(*Response, error) {
	fmt.Println("hello")
}
