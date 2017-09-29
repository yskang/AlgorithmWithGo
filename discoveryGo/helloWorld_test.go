package discoveryGo

import (
	"net/http"
	"fmt"
	"log"
)

func serve_main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, 세계!")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ExampleServeMain() {
	serve_main()
	// output:
	//
}
