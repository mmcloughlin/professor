package professor

import (
	"log"
	"net/http"
	"net/http/pprof"
)

func init() {
	http.DefaultServeMux = http.NewServeMux()
}

func Handle(mux *http.ServeMux) {
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
}

func NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	Handle(mux)
	return mux
}

func NewServer(addr string) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: NewServeMux(),
	}
}

func ListenAndServe(addr string) error {
	return NewServer(addr).ListenAndServe()
}

func Launch(addr string) {
	go func() {
		log.Fatal(ListenAndServe(addr))
	}()
}
