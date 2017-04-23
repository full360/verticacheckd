package checkd

import (
	"fmt"
	"os/exec"
	"regexp"
)

type CheckService interface {
	Health() (bool, error)
}

type checkService struct {
	address string
	cmd     string
	cmdArgs []string
}

func NewCheckService(addr, cmd string, cmdArgs []string) *checkService {
	return &checkService{
		address: addr,
		cmd:     cmd,
		cmdArgs: cmdArgs,
	}
}

func (c checkService) Health() (bool, error) {
	cmdOut, err := exec.Command(c.cmd, c.cmdArgs...).CombinedOutput()
	if err != nil {
		return false, err
	}

	exp := regexp.MustCompile(fmt.Sprintf(`.%s\s+\W\s+UP`, c.address))
	return exp.Match(cmdOut), nil
}
