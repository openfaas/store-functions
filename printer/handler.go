package function

import (
	"fmt"
	"io"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var input []byte

	if r.Body != nil {
		defer r.Body.Close()

		body, _ := io.ReadAll(r.Body)

		input = body
	}

	for k, v := range r.Header {
		fmt.Printf("%s=%v\n", k, v)
	}

	fmt.Println()
	fmt.Println(string(input))
	fmt.Println()

	w.WriteHeader(http.StatusAccepted)
}
