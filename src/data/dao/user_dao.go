package dao

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

type userDAO struct {
	executor contract.Executor
}

func NewUserDAO(executor contract.Executor) contract.UserDAO {
	return &userDAO{
		executor: executor,
	}
}

func (u *userDAO) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	ID, err := u.getUserID(username)
	if err != nil {
		return nil, err
	}

	expiresAt, err := u.getExpirationDate(ctx, username)
	if err != nil {
		return nil, err
	}

	limit := u.getConnectionLimit(ctx, username)

	user := &entity.User{
		ID:        ID,
		Username:  username,
		ExpiresAt: expiresAt,
		Limit:     limit,
	}
	return user, nil
}

func (u *userDAO) getConnectionLimit(ctx context.Context, username string) int {
	connLimitPattern := regexp.MustCompile(`connection_limit:\s*(\d+)`)
	phpLimitPattern := regexp.MustCompile(`\|\s*(\d+)`)
	archivePattern := regexp.MustCompile(fmt.Sprintf(`%s\s+(\d+)`, username))

	vpsOut, _ := u.executeCommand(ctx, fmt.Sprintf("vps view -u %s", username))
	vpsMatches := connLimitPattern.FindAllStringSubmatch(vpsOut, -1)

	limit := 1
	if len(vpsMatches) > 0 {
		if n, err := strconv.Atoi(vpsMatches[0][1]); err == nil {
			limit = n
		}
	}

	phpOut, _ := u.executeCommand(ctx, fmt.Sprintf("php /opt/DragonCore/menu.php printlim2 %s", username))
	phpMatches := phpLimitPattern.FindAllStringSubmatch(phpOut, -1)

	if len(phpMatches) > 0 {
		if n, err := strconv.Atoi(phpMatches[0][1]); err == nil {
			limit = n
		}
	}

	data, err := os.ReadFile("/root/usuarios.db")
	if err == nil {
		if archMatches := archivePattern.FindStringSubmatch(string(data)); len(archMatches) > 1 {
			if n, err := strconv.Atoi(archMatches[1]); err == nil {
				limit = n
			}
		}
	}

	return limit
}

func (u *userDAO) getExpirationDate(ctx context.Context, username string) (time.Time, error) {
	command := fmt.Sprintf("chage -l %s", username)
	output, err := u.executor.Execute(ctx, command)
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

func (u *userDAO) getUserID(username string) (int, error) {
	data, err := os.ReadFile("/etc/passwd")
	if err != nil {
		return -1, err
	}

	pattern := regexp.MustCompile(fmt.Sprintf(`%s\:.*?\:(\d+)\:.*`, username))
	matches := pattern.FindStringSubmatch(string(data))
	if len(matches) < 2 {
		return -1, nil
	}

	return strconv.Atoi(matches[1])
}

func (u *userDAO) executeCommand(ctx context.Context, command string) (string, error) {
	output, err := u.executor.Execute(ctx, command)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(output), nil
}
