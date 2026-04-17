package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Run app in port :", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "guest"
		}
		dataDir := os.Getenv("APP_DATA")
		if dataDir == "" {
			dataDir = "/logs"
		}
		filename := fmt.Sprintf("%s/%s.txt", dataDir, name)
		content := fmt.Sprintf("Hello %s!\n", name)
		err := os.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			fmt.Fprintf(w, "Error: %v\n", err)
			return
		}
		fmt.Fprintf(w, "DONE Write File : %s\n", filename)
		fmt.Printf("DONE Write File : %s\n", filename)
	})

	http.ListenAndServe(":"+port, nil)
}
