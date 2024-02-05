package data

import (
	"bytes"
	"context"
	"os/exec"
	"strings"

	"github.com/DTunnel0/CheckUser-Go/src/domain/contract"
)

type bashExecutor struct {
	pool chan *exec.Cmd
}

func NewBashExecutor() contract.Executor {
	poolSize := 100
	pool := make(chan *exec.Cmd, poolSize)
	for i := 0; i < poolSize; i++ {
		pool <- exec.Command("bash")
	}
	return &bashExecutor{pool: pool}
}

func (b *bashExecutor) Execute(ctx context.Context, command string) (string, error) {
	cmd := <-b.pool
	defer func() {
		b.pool <- cmd
	}()

	cmd.Args = []string{"bash", "-c", command}
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}
