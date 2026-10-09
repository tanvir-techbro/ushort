package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func startServer() error {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("Page")))
	mux.HandleFunc("POST /shorten", func(w http.ResponseWriter, r *http.Request) {
		var v struct {
			URL string `json:"url"`
		}
		if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
			http.Error(w, "bad json", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"url": v.URL})
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
