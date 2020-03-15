package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: gocurl <url>")
		os.Exit(1)
	}
}

func main() {
	h, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	j, _ := json.Marshal(h)
	fmt.Println(j)
}
