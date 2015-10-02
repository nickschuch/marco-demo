package main

import (
	"github.com/nickschuch/marco-lib"
)

const name = "demo"

func main() {
	b := []marco.Backend{
		marco.Backend{
			Type:   name,
			Domain: "example.com",
			List: []string{
				"http://localhost:8080",
			},
		},
	}
	err := marco.Send(b, "http://localhost:81")
	if err != nil {
		panic(err)
	}
}
