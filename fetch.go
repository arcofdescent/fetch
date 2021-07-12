package main

import (
	"fmt"
    "net/http"
    "os"
)

func main() {
    url := "http://app.edyst.com"

	fmt.Println("fetch a URL")
    resp, err = http.Get(url)
    if err != nil {
        fmf.Printf("Got error: %v\n", err)
        os.Exit(1)
    }
}
