package todo_list

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func MainFn() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8080", nil)
}
