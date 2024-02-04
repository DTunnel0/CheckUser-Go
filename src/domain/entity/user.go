package entity

import "time"

type User struct {
	ID        int
	Username  string
	ExpiresAt time.Time
	Limit     int
}

func (u *User) LimitReached(devices int) bool {
	return u.Limit <= devices
}
