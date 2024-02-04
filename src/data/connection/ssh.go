package connection

import (
	"context"
	"fmt"
	"strconv"

	"github.com/DTunnel0/CheckUser-Go/src/domain/contract"
)

type sshConnection struct {
	executor contract.Executor
	next     contract.Connection
}

func NewSSHConnection(executor contract.Executor) contract.Connection {
	return &sshConnection{executor: executor}
}

func (ssh *sshConnection) SetNext(connection contract.Connection) {
	ssh.next = connection
}

func (s *sshConnection) CountByUsername(ctx context.Context, username string) int {
	cmd := fmt.Sprintf("ps -u %s | grep -c '[s]shd'", username)
	result, _ := s.executor.Execute(ctx, cmd)
	count, err := strconv.Atoi(result)

	totalConnections := 0
	if err == nil {
		totalConnections += count
	}

	if s.next != nil {
		totalConnections += s.next.CountByUsername(ctx, username)
	}

	return totalConnections
}

func (s *sshConnection) Count(ctx context.Context) int {
	cmd := fmt.Sprint("ps -ef | grep '[s]shd' | grep -cEv 'root|nobody|grep'")
	result, _ := s.executor.Execute(ctx, cmd)
	count, err := strconv.Atoi(result)
	if err != nil {
		return 0
	}
	return count - 1
}
