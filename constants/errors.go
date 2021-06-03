package constants

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidBoardSize    = fmt.Errorf("board size should be greater or equal than %d", MinBoardSize)
	ErrTargetIsUnreachable = errors.New("the target is unreachable")
	ErrPieceDoesNotExist   = errors.New("piece does not exist")
)
