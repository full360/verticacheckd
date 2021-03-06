package verticacheckd

import (
	"reflect"
	"testing"
)

func TestService_NewService(t *testing.T) {
	data := struct {
		address string
		name    string
		args    []string
	}{
		address: "10.0.1.66",
		name:    "cat",
		args:    []string{"testdata/sample_output.txt"},
	}

	cases := []struct {
		svc *checkService
	}{
		{
			&checkService{
				address: data.address,
				cmd:     data.name,
				cmdArgs: data.args,
			},
		},
	}

	for _, c := range cases {
		svc := NewService(data.address, data.name, data.args)

		if !reflect.DeepEqual(svc, c.svc) {
			t.Errorf("expected %v to be %v", svc, c.svc)
		}
	}
}

func TestService_HostState(t *testing.T) {
	command := struct {
		name string
		args []string
	}{
		name: "cat",
		args: []string{"testdata/sample_output.txt"},
	}

	cases := []struct {
		found bool
		svc   *checkService
	}{
		{
			true,
			&checkService{
				address: "10.0.1.66",
				cmd:     command.name,
				cmdArgs: command.args,
			},
		},
		{
			true,
			&checkService{
				address: "172.31.47.139",
				cmd:     command.name,
				cmdArgs: command.args,
			},
		},
		{
			false,
			&checkService{
				address: "172.31.47.100",
				cmd:     command.name,
				cmdArgs: command.args,
			},
		},
	}

	for _, c := range cases {
		check, err := c.svc.HostState()
		if err != nil {
			t.Errorf("expected %v to be %v", nil, err)
		}

		if check != c.found {
			t.Errorf("expected %v to be %v", c.found, check)
		}
	}
}

func TestService_DBHostState(t *testing.T) {
	command := struct {
		name string
		args []string
	}{
		name: "cat",
		args: []string{"testdata/sample_output.txt"},
	}

	cases := []struct {
		db    string
		found bool
		svc   *checkService
	}{
		{
			"climatedb",
			true,
			&checkService{
				address: "10.0.1.66",
				cmd:     command.name,
				cmdArgs: command.args,
			},
		},
		{
			"climatedb",
			true,
			&checkService{
				address: "172.31.47.139",
				cmd:     command.name,
				cmdArgs: command.args,
			},
		},
		{
			"climatedb",
			false,
			&checkService{
				address: "172.31.47.100",
				cmd:     command.name,
				cmdArgs: command.args,
			},
		},
		{
			"otherdb",
			false,
			&checkService{
				address: "10.0.1.66",
				cmd:     command.name,
				cmdArgs: command.args,
			},
		},
	}

	for _, c := range cases {
		check, err := c.svc.DBHostState(c.db)
		if err != nil {
			t.Errorf("expected %v to be %v", nil, err)
		}

		if check != c.found {
			t.Errorf("expected %v to be %v", c.found, check)
		}
	}
}
