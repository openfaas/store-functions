// Copyright (c) 2023 Alex Ellis, OpenFaaS Ltd
// Licensed under the MIT license, see LICENSE.md file.

package function

import (
	"io"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {

	if r.Body != nil {
		defer r.Body.Close()
	}

	res, err := http.Get("https://api.ipify.org")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for k, v := range res.Header {
		w.Header().Set(k, v[0])
	}

	if res.Body != nil {
		defer res.Body.Close()
		io.Copy(w, res.Body)
	}
}
