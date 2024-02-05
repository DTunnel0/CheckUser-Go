package data

import (
	"context"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/DTunnel0/CheckUser-Go/src/domain/contract"
)

type bashExecutor struct {
	cache map[string]*cacheItem
	mutex sync.Mutex
}

type cacheItem struct {
	expiry time.Time
	value  string
}

func NewBashExecutor() contract.Executor {
	return &bashExecutor{
		cache: make(map[string]*cacheItem),
	}
}

func (b *bashExecutor) getCachedResult(command string) (string, bool) {
	b.mutex.Lock()
	cachedItem, ok := b.cache[command]
	b.mutex.Unlock()

	if ok && time.Now().Before(cachedItem.expiry) {
		return cachedItem.value, true
	}

	return "", false
}

func (b *bashExecutor) Execute(ctx context.Context, command string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	cachedResult, ok := b.getCachedResult(command)
	if ok {
		return cachedResult, nil
	}

	cmd := exec.CommandContext(ctx, "bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	result := strings.TrimSpace(string(output))

	b.mutex.Lock()
	b.cache[command] = &cacheItem{
		expiry: time.Now().Add(time.Second * 30),
		value:  result,
	}
	b.mutex.Unlock()

	return result, nil
}
