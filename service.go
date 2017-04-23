package checkd

import (
	"fmt"
	"os/exec"
	"regexp"
)

type CheckService interface {
	HostState() (bool, error)
}

type checkService struct {
	address string
	cmd     string
	cmdArgs []string
}

func NewService(addr, cmd string, cmdArgs []string) *checkService {
	return &checkService{
		address: addr,
		cmd:     cmd,
		cmdArgs: cmdArgs,
	}
}

func (c checkService) state(regex string) (bool, error) {
	cmdOut, err := exec.Command(c.cmd, c.cmdArgs...).CombinedOutput()
	if err != nil {
		return false, err
	}

	exp := regexp.MustCompile(regex)
	return exp.Match(cmdOut), nil
}

func (c checkService) HostState() (bool, error) {
	return c.state(fmt.Sprintf(`.%s\s+\W\s+UP`, c.address))
}
