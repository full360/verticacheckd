package verticacheckd

import (
	"fmt"
	"net/http"
	"os/exec"
	"regexp"
)

func Check(addr, cmd string, args ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cmdOut, err := exec.Command(cmd, args...).CombinedOutput()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error running the command"))
		}

		rgxp := fmt.Sprintf(`.%s\s+\W\s+UP`, addr)
		exp := regexp.MustCompile(rgxp)

		match := exp.Match(cmdOut)
		if match {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}
