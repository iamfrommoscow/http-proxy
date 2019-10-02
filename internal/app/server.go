package server

import (
	"crypto/tls"
	"log"
	"net/http"
	"proxy/internal/pkg/db"
	"proxy/internal/pkg/handlers"
)

func StartApp(port string) {
	db.Connect()

	server := &http.Server{
		Addr: ":" + port,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodConnect {
				handlers.HandleTunneling(w, r)
			} else {
				handlers.HandleHTTP(w, r)
			}
		}),

		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	log.Fatal(server.ListenAndServe())

}
