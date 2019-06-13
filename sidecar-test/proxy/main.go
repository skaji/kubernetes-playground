package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	proxyHost := "http://localhost:9000"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		proxyURL := proxyHost + path
		res, err := http.Get(proxyURL)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		defer res.Body.Close()
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		fmt.Fprint(w, string(b))
	})
	http.ListenAndServe(":8080", nil)
}
