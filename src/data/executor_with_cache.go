package data

import (
	"context"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/DTunnel0/CheckUser-Go/src/domain/contract"
)

const cacheExpirationTime = 5 * time.Minute

type bashExecutorWithCache struct {
	cache sync.Map
}

type cachedResult struct {
	result      string
	lastFetched time.Time
}

func NewBashExecutorWithCache() contract.Executor {
	return &bashExecutorWithCache{}
}

func (b *bashExecutorWithCache) Execute(ctx context.Context, command string) (string, error) {
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

func (b *bashExecutorWithCache) getCachedResult(command string) (string, bool) {
	cached, ok := b.cache.Load(command)
	if !ok {
		return "", false
	}

	cachedResult, ok := cached.(cachedResult)
	if !ok || time.Since(cachedResult.lastFetched) >= cacheExpirationTime {
		return "", false
	}

	return cachedResult.result, true
}

func (b *bashExecutorWithCache) executeCommand(ctx context.Context, command string) (string, error) {
	args := strings.Split(command, " ")
	cmd := exec.CommandContext(ctx, args[0], args[1:]...)

	resultBytes, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(resultBytes)), nil
}

func (b *bashExecutorWithCache) cacheResult(command, result string) {
	b.cache.Store(command, cachedResult{
		result:      result,
		lastFetched: time.Now(),
	})
}
