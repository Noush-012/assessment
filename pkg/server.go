package pkg

import (
	"fmt"
	"net/http"

	"github.com/noush/pkg/handler"
)

func Server() {

	http.HandleFunc("/v1", handler.PostHandler)

	port := ":7001"
	fmt.Printf("\nServer is listening on port %s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Failed to start server", err)
	}
}
