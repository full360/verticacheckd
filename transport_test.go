package verticacheckd

import (
	"errors"
	"fmt"
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

func (c stubCheckService) DBHostState(db string) (bool, error) {
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

func TestTransport_DBStateHandler(t *testing.T) {
	cases := []struct {
		expectedStatus int
		name           string
		svc            *stubCheckService
	}{
		{
			http.StatusOK,
			"somedb",
			&stubCheckService{true, nil},
		},
		{
			http.StatusInternalServerError,
			"otherdb",
			&stubCheckService{false, nil},
		},
		{
			http.StatusInternalServerError,
			"someotherdb",
			&stubCheckService{false, errors.New("some error")},
		},
	}

	for _, c := range cases {
		path := fmt.Sprintf("/dbs/%s/state", c.name)

		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		handler := DBStateHandler(c.svc)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != c.expectedStatus {
			t.Errorf("expected %v to be %v", status, c.expectedStatus)
		}
	}
}
