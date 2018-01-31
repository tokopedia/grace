package grace

import (
	"log"
	"net/http"
	"time"
)

func ExampleServe() {
	http.HandleFunc("/foo/bar", foobarHandler)
	cfg := &Config{
		Timeout:      5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	// log.Fatal(grace.Serve(":9000", nil, cfg))
	log.Fatal(Serve(":9000", nil, cfg))

}

func foobarHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foobar"))
}
