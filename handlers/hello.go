package handlers

import "net/http"
import "log"
import "io/ioutil"
import "fmt"

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello world!")
	d, err := ioutil.ReadAll(r.Body) 

	if err != nil {
		http.Error(rw, "Ooops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Data %s\n", d)
}