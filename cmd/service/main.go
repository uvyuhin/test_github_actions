package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/uvyuhin/test_github_actions/handlers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf("Hello World! Result: %v", handlers.SomeNumber(rand.Intn(100))))
	})
	_ = http.ListenAndServe("localhost:8089", nil)

}
