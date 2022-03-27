package domain

import (
	"errors"
	"fmt"
)

var (
	InternalServerErr = errors.New("Internal server error")
)

func UserNotFoundErr(user User) error {
	return fmt.Errorf("User id = %d not found", user.ID)
}
