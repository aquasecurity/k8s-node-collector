package collector

import (
	"os/exec"
	"strings"
)

const (
	// WorkerNode worker node type
	WorkerNode = "worker"
	// MasterNode master Node type
	MasterNode   = "master"
	shellCommand = "sh"
)

var (
	replacments = map[string]string{
		"\n":         ",",
		"[^\"]\\S*'": "",
	}
)

// Shell command interface to preform shell exec commands
type Shell interface {
	Execute(commandArgs string) (string, error)
	FindNodeType() (string, error)
}

// NewShellCmd instansiate new shell command
func NewShellCmd() Shell {
	return &cmd{}
}

type cmd struct {
}

// Execute execute a shell command and retun it output or error
func (e *cmd) Execute(commandArgs string) (string, error) {
	cm := exec.Command(shellCommand, "-c", commandArgs)
	output, err := cm.CombinedOutput()
	if err != nil {
		return "", nil
	}
	// trim newline
	outPutWithDelimiter := SanitizeString(strings.TrimSuffix(string(output), "\n"), replacments)
	return outPutWithDelimiter, nil
}

func (e *cmd) FindNodeType() (string, error) {
	masterConfigFiles := []string{
		"ls /etc/kubernetes/controller-manager.conf",
		"ls /etc/kubernetes/manifests/kube-apiserver.yaml",
		"ls /etc/kubernetes/scheduler.conf",
	}
	for _, path := range masterConfigFiles {
		output, err := e.Execute(path)
		if err != nil {
			return WorkerNode, nil
		}
		outputParts := strings.Split(output, ",")
		if len(outputParts) > 0 {
			for _, part := range outputParts {
				if (len(strings.TrimSpace(part)) != 0 && !strings.Contains(path, strings.TrimSpace(part))) ||
					len(strings.TrimSpace(part)) == 0 ||
					err != nil {
					return WorkerNode, nil
				}
			}
		}
	}
	return MasterNode, nil
}
