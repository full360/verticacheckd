package checkd

import (
	"net/http"
)

func StateHandler(svc CheckService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		check, err := svc.HostState()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error"))
		}

		if check {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}
