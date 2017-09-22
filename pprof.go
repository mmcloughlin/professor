package professor

import (
	"log"
	"net/http"
	"net/http/pprof"
)

// init disables default handlers registered by importing net/http/pprof.
func init() {
	http.DefaultServeMux = http.NewServeMux()
}

// Handle adds standard pprof handlers to mux.
func Handle(mux *http.ServeMux) {
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
}

// NewServeMux builds a ServeMux and populates it with standard pprof handlers.
func NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	Handle(mux)
	return mux
}

// NewServer constructs a server at addr with the standard pprof handlers.
func NewServer(addr string) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: NewServeMux(),
	}
}

// ListenAndServe starts a server at addr with standard pprof handlers.
func ListenAndServe(addr string) error {
	return NewServer(addr).ListenAndServe()
}

// Launch a standard pprof server at addr.
func Launch(addr string) {
	go func() {
		log.Fatal(ListenAndServe(addr))
	}()
}
