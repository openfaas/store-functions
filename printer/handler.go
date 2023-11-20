package function

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func formatJSON(data []byte) ([]byte, error) {
	var out bytes.Buffer

	if err := json.Indent(&out, data, "", "    "); err != nil {
		return data, err
	}

	return out.Bytes(), nil
}

func printRaw() bool {
	val := os.Getenv("RAW")
	raw, err := strconv.ParseBool(val)
	if err != nil {
		return false
	}
	return raw
}

func Handle(w http.ResponseWriter, r *http.Request) {
	var input []byte

	contentType := r.Header.Get("content-type")
	if r.Body != nil {
		defer r.Body.Close()

		body, _ := io.ReadAll(r.Body)

		if !printRaw() && contentType == "application/json" {
			var err error
			input, err = formatJSON(body)
			if err != nil {
				http.Error(w, fmt.Sprintf("failed to format JSON body: %s", err), http.StatusInternalServerError)
				return
			}
		} else {
			input = body
		}
	}

	for k, v := range r.Header {
		fmt.Printf("%s=%v\n", k, v)
	}

	fmt.Println()
	fmt.Println(string(input))
	fmt.Println()

	w.WriteHeader(http.StatusAccepted)
}
