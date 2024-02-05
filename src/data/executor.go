package data

import (
	"context"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/DTunnel0/CheckUser-Go/src/domain/contract"
)

const cacheExpiresTime = 1 * time.Minute

type bashExecutor struct {
	cache sync.Map
}

type cachedResult struct {
	result      string
	lastFetched time.Time
}

func NewBashExecutor() contract.Executor {
	return &bashExecutor{}
}

func (b *bashExecutor) Execute(ctx context.Context, command string) (string, error) {
	if cachedResult, ok := b.getCachedResult(command); ok {
		return cachedResult, nil
	}

	result, err := b.executeCommand(ctx, command)
	if err != nil {
		return "", err
	}

	b.cacheResult(command, result)
	return result, nil
}

func (b *bashExecutor) getCachedResult(command string) (string, bool) {
	if cached, ok := b.cache.Load(command); ok {
		if cachedResult, ok := cached.(cachedResult); ok && time.Since(cachedResult.lastFetched) < cacheExpiresTime {
			return cachedResult.result, true
		}
	}
	return "", false
}

func (b *bashExecutor) executeCommand(ctx context.Context, command string) (string, error) {
	args := strings.Split(command, " ")
	cmd := exec.CommandContext(ctx, args[0], args[1:]...)

	resultBytes, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(resultBytes)), nil
}

func (b *bashExecutor) cacheResult(command, result string) {
	b.cache.Store(command, cachedResult{
		result:      result,
		lastFetched: time.Now(),
	})
}
