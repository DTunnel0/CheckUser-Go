package connection

import (
	"context"
	"fmt"
	"regexp"

	"github.com/DTunnel0/CheckUser-Go/src/domain/contract"
)

type sshConnection struct {
	executor contract.Executor
	next     contract.CountConnection
}

func NewSSHConnection(executor contract.Executor) contract.CountConnection {
	return &sshConnection{executor: executor}
}

func (ssh *sshConnection) SetNext(connection contract.CountConnection) {
	ssh.next = connection
}

func (s *sshConnection) ByUsername(ctx context.Context, username string) (int, error) {
	cmd := "ps -u " + username
	result, err := s.executor.Execute(ctx, cmd)
	if err != nil {
		return 0, err
	}

	sshdPattern := regexp.MustCompile(`.*sshd`)
	matches := sshdPattern.FindAllStringSubmatch(result, -1)
	totalConnections := len(matches)
	if s.next != nil {
		count, err := s.next.ByUsername(ctx, username)
		if err == nil {
			totalConnections += count
		}
	}

	return totalConnections, err
}

func (s *sshConnection) All(ctx context.Context) (int, error) {
	cmd := "ps -ef"
	result, err := s.executor.Execute(ctx, cmd)
	if err != nil {
		return 0, fmt.Errorf("failed to execute command: %w", err)
	}

	sshdPattern := regexp.MustCompile(`(?m)^(\S+)\s+\d+\s+\d+\s+\d+\s+\d+:\d+\s+.*\bsshd\b.*$`)
	processMatches := sshdPattern.FindAllStringSubmatch(string(result), -1)
	forbiddenUsernames := map[string]bool{
		"root":   true,
		"nobody": true,
		"grep":   true,
	}

	totalConnections := 0
	for _, match := range processMatches {
		username := match[1]
		if !forbiddenUsernames[username] {
			totalConnections++
		}
	}

	if s.next != nil {
		count, err := s.next.All(ctx)
		if err == nil {
			totalConnections += count
		}
	}

	return totalConnections, nil
}
