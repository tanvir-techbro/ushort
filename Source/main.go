package main

import (
	"fmt"
	"net/http"
	"os"
)

func startServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	})

	fmt.Println("Listening on :8080")
	return http.ListenAndServe(":8080", mux)
}

func main() {
	if len(os.Args) > 1 {
		if err := parseCli(os.Args[1:]); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("No commands given.")
		fmt.Println("Run --help, -h for more info.")
	}
}
