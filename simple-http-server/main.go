package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi,seven\n"))
	})
	http.ListenAndServe(":17778", nil)
}
