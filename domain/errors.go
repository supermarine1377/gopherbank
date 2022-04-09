package domain

import (
	"errors"
	"fmt"
)

var (
	InternalServerErr = errors.New("Internal server error")
)

func UserNotFoundErr(userId int) error {
	return fmt.Errorf("User id = %d not found", userId)
}
