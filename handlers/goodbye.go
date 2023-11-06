package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Bye struct {
	l *log.Logger
}

func NewBye(l *log.Logger) *Bye {
	return &Bye{l}
}

func (b *Bye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Something wrong!", http.StatusBadRequest)
		return
	}

	b.l.Printf("Request at `/bye` with data : `%s`", data)

	fmt.Fprintf(rw, "Bye %s", data)

}
