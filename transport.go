package verticacheckd

import (
	"net/http"

	"github.com/gorilla/mux"
)

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
