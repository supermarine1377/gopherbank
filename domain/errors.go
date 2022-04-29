package domain

import (
	"errors"
	"fmt"
)

// インフラストラクチャ層の関数が返すエラー
var (
	ErrInternalServer = errors.New("Internal server error")
)

// インターフェース層の関数が返すエラー
var (
	ErrInvalidUserCreateReq = errors.New("Invalid user registration request")
)

func ErrUserNotFound(userId int) error {
	return fmt.Errorf("User id = %d not found", userId)
}
