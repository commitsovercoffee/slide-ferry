package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	// http handlers ~ handles request for specified path. ----------------

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Something wrong!", http.StatusBadRequest)
			return
		}

		log.Printf("Received request with data : `%s`", data)

		fmt.Fprintf(rw, "Hello %s", data)
	})

	http.HandleFunc("/bye", func(rw http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(rw, "See ya!")

	})

	// http server
	http.ListenAndServe(":9090", nil)

}
