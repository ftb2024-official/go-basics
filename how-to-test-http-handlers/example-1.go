package howtotesthttphandlers

import "net/http"

func Example1(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	rw.WriteHeader(http.StatusOK)
	_, _ = rw.Write([]byte("FOO"))
}
