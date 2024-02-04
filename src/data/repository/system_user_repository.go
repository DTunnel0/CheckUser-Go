package repository

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/DTunnel0/CheckUser-Go/src/domain/contract"
	"github.com/DTunnel0/CheckUser-Go/src/domain/entity"
)

const (
	idCommand    = "id -u"
	chageCommand = "chage -l"
	vpsCommand   = "vps view -u"
	archivePath  = "/root/usuarios.db"
)

type systemUserRepository struct {
	executor contract.Executor
}

func NewSystemUserRepository(executor contract.Executor) contract.UserRepository {
	return &systemUserRepository{
		executor: executor,
	}
}

func (r *systemUserRepository) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	id, err := r.executeCommand(ctx, fmt.Sprintf("%s %s", idCommand, username))
	if err != nil {
		return nil, err
	}

	ID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	expiresAt, err := r.getExpirationDate(ctx, username)
	if err != nil {
		return nil, err
	}

	limit := r.getConnectionLimit(ctx, username)

	return &entity.User{
		ID:        ID,
		Username:  username,
		ExpiresAt: expiresAt,
		Limit:     limit,
	}, nil
}

func (r *systemUserRepository) getConnectionLimit(ctx context.Context, username string) int {
	limit := 1

	cmd := fmt.Sprintf("%s %s | grep connection_limit: | cut -d' ' -f2", vpsCommand, username)
	output, err := r.executeCommand(ctx, cmd)
	if err == nil {
		num, err := strconv.Atoi(output)
		if err == nil {
			limit = num
		}
	}

	data, err := os.ReadFile(archivePath)
	if err == nil {
		re := regexp.MustCompile(fmt.Sprintf("%s\\s+(\\d+)", username))
		match := re.FindStringSubmatch(string(data))
		if len(match) > 1 {
			num, err := strconv.Atoi(match[1])
			if err == nil {
				limit = num
			}
		}
	}

	return limit
}

func (r *systemUserRepository) getExpirationDate(ctx context.Context, username string) (time.Time, error) {
	command := fmt.Sprintf("chage -l %s", username)
	output, err := r.executor.Execute(ctx, command)
	if err != nil {
		return time.Time{}, err
	}

	search := regexp.MustCompile(`Account expires\s*:\s*(.*)`).FindStringSubmatch(output)
	if len(search) < 2 {
		return time.Time{}, nil
	}

	expirationDate, err := time.Parse("Jan 02, 2006", search[1])
	if err != nil {
		return time.Time{}, err
	}

	return expirationDate, nil
}

func (r *systemUserRepository) executeCommand(ctx context.Context, command string) (string, error) {
	output, err := r.executor.Execute(ctx, command)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(output), nil
}
