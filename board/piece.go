package board

import (
	"github.com/felipeap92/the-knights-travails/constants"
)

type Piece interface {
	GetValidMovements() []Vector2D
}

func CreateNewPiece(pieceName string) (Piece, error) {
	if pieceName == "knight" {
		return &Knight{}, nil
	}

	if pieceName == "camel" {
		return &Camel{}, nil
	}

	return nil, constants.ErrPieceDoesNotExist
}

type Knight struct{}

func (knight *Knight) GetValidMovements() []Vector2D {
	return []Vector2D{{1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}}
}

type Camel struct{}

func (camel *Camel) GetValidMovements() []Vector2D {
	return []Vector2D{{1, 3}, {3, 1}, {3, -1}, {1, -3}, {-1, -3}, {-3, -1}, {-3, 1}, {-1, 3}}
}
