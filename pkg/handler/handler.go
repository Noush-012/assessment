package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/noush/pkg/service"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "Request should be in POST method")
		return
	}

	reqData, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Unable to read request data")
		return
	}

	rqstChannel := make(chan []byte)

	jsonData, err := json.Marshal(reqData)
	if err != nil {
		fmt.Println("failed to marshal request data")
		return
	}

	rqstChannel <- jsonData

	go service.GoWorker(rqstChannel)

	fmt.Println(string(reqData))
	w.WriteHeader(200)

}
