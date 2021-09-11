package main

import "net/http"
import "log"
import "io/ioutil"
import "fmt"

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request) {
		log.Println("Hello world!")
		d, err := ioutil.ReadAll(r.Body) 

		if err != nil {
			http.Error(rw, "Ooops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Data %s\n", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye world!")
	})

	http.ListenAndServe(":9090", nil)
}