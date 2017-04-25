package verticacheckd

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func AddLogger(l *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		l.Println(r.Method, r.URL.Path, time.Since(start))
	})
}

func StateHandler(svc CheckService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		check, err := svc.HostState()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			if check {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	})
}

func DBStateHandler(svc CheckService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		check, err := svc.DBHostState(vars["name"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			if check {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	})
}
