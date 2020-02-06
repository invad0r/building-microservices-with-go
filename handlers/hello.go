package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// satisfies http-handler-interface
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oooops", http.StatusBadRequest)
		//rw.WriteHeader(http.StatusBadRequest)
		//rw.Write([]byte("Oooops"))
		return
	}

	log.Printf("Data: %s\n", d)
	fmt.Fprintf(rw, "hi %s\n", d)
}
