package testhelper

import (
	"net/http"
	"net/http/httptest"
)

// NewServer :
func NewServer(opts ...func(*http.ServeMux)) (*httptest.Server, func()) {
	mux := http.NewServeMux()
	for _, opt := range opts {
		opt(mux)
	}
	ts := httptest.NewServer(mux)
	return ts, ts.Close
}

// WithHandlerOption :
func WithHandlerOption(
	path, method string,
	status int,
	respBody []byte,
) func(*http.ServeMux) {
	return func(mux *http.ServeMux) {
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			if r.Method == method {
				w.WriteHeader(status)
				_, _ = w.Write(respBody)
			}
		})
	}
}
