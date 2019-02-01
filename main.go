package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func f(src string) string {
	return fmt.Sprintf("Hello, %s", src)
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res := f(html.EscapeString(r.URL.Path))
		fmt.Fprint(w, res)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
