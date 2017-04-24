package verticacheckd

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type stubCheckService struct {
	resp bool
	err  error
}

func (c stubCheckService) HostState() (bool, error) {
	return c.resp, c.err
}

func TestTransport_StateHandler(t *testing.T) {
	cases := []struct {
		expectedStatus int
		svc            *stubCheckService
	}{
		{
			http.StatusOK,
			&stubCheckService{true, nil},
		},
		{
			http.StatusInternalServerError,
			&stubCheckService{false, nil},
		},
		{
			http.StatusInternalServerError,
			&stubCheckService{false, errors.New("some error")},
		},
	}

	req, err := http.NewRequest("GET", "/state", nil)
	if err != nil {
		t.Fatal(err)
	}

	for _, c := range cases {
		rr := httptest.NewRecorder()

		handler := StateHandler(c.svc)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != c.expectedStatus {
			t.Errorf("expected %v to be %v", status, c.expectedStatus)
		}
	}
}
