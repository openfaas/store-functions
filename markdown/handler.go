package function

import (
	"io"
	"net/http"

	"github.com/russross/blackfriday/v2"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var input []byte

	if r.Body != nil {
		defer r.Body.Close()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read body", http.StatusInternalServerError)
			return
		}

		input = body
	}

	// Convert Markdown to HTML
	htmlContent := blackfriday.Run(input)

	// Set response header for HTML
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(htmlContent)
}
