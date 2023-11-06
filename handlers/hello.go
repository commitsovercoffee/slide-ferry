package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Something wrong!", http.StatusBadRequest)
		return
	}

	h.l.Printf("Request at `/` with data : `%s`", data)

	fmt.Fprintf(rw, "Hello %s", data)

}
