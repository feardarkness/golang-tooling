// demo is an awesome demo
package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8087", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// var re = regexp.MustCompile("^/(.+)@golang.org$")

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	// re := regexp.MustCompile("^/(.+)@golang.org$")
	// path := r.URL.Path
	path := r.URL.Path[1:]
	// match := re.FindAllStringSubmatch(path, -1)
	// if match != nil {
	if strings.HasSuffix(path, "@golang.org") {
		fmt.Fprintf(w, "Hello gopher %s", strings.TrimSuffix(path, "@golang.org"))
		return
	}
	fmt.Fprintf(w, "Hello dear %s", path)
}
