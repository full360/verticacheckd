package verticacheckd

import (
	"fmt"
	"os/exec"
	"regexp"
)

type CheckService interface {
	HostState() (bool, error)
	DBHostState(db string) (bool, error)
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

func (c checkService) DBHostState(db string) (bool, error) {
	return c.state(fmt.Sprintf(`\s+%s\s+\W\s+%s\s+\W\s+UP`, db, c.address))
}
