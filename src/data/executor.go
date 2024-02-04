package data

import (
	"context"
	"os/exec"
	"strings"

	"github.com/DTunnel0/CheckUser-Go/src/domain/contract"
)

type bashExecutor struct {
}

func NewBashExecutor() contract.Executor {
	return &bashExecutor{}
}

func (b *bashExecutor) Execute(ctx context.Context, command string) (string, error) {
	cmd := exec.CommandContext(ctx, "bash", "-c", command)
	result, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(result)), nil
}
