package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	dir := "./data"

	fs := http.FileServer(http.Dir(dir))

	http.Handle("/", cors(fs))

	// http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(700)
	// 	w.Write([]byte("{\"code\":700}"))
	// })

	fmt.Println("http://localhost:4560")

	log.Fatal(http.ListenAndServe(":4560", nil))
}
func cors(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// do your cors stuff
		// return if you do not want the FileServer handle a specific request
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		fs.ServeHTTP(w, r)
	}
}
