package connection

import (
	"context"
	"fmt"
	"regexp"

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

func (s *sshConnection) CountByUsername(ctx context.Context, username string) (int, error) {
	cmd := "ps -u " + username
	result, err := s.executor.Execute(ctx, cmd)
	if err != nil {
		return 0, err
	}
	sshdPattern := regexp.MustCompile(`(?m)^(\S+)\s+\d+\s+\d+\s+\d+\s+\d+:\d+\s+.*\bsshd\b.*$`)
	matches := sshdPattern.FindAllStringSubmatch(result, -1)
	totalConnections := len(matches)

	if s.next != nil {
		count, err := s.next.CountByUsername(ctx, username)
		if err == nil {
			totalConnections += count
		}
	}

	return totalConnections, err
}
func (s *sshConnection) Count(ctx context.Context) (int, error) {
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
		count, err := s.next.Count(ctx)
		if err == nil {
			totalConnections += count
		}
	}

	return totalConnections, nil
}
