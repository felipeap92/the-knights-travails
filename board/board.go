package board

import (
	"github.com/felipeap92/the-knights-travails/constants"
)

type Vector2D struct {
	X int
	Y int
}

type PathNode struct {
	Parent *PathNode
	Pos    Vector2D
}

type Board struct {
	size int
}

func NewBoard(size int) (*Board, error) {
	if size < constants.MinBoardSize {
		return nil, constants.ErrInvalidBoardSize
	}

	return &Board{size}, nil
}

func (board *Board) FindShortestPath(startPos, targetPos Vector2D) ([]Vector2D, error) {
	validKnightMovements := []Vector2D{{1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}}
	queue := []*PathNode{{
		Parent: nil,
		Pos:    startPos,
	}}
	isCellVisited := make([][]bool, board.size)
	for i := 0; i < board.size; i++ {
		isCellVisited[i] = make([]bool, board.size)
	}

	var node *PathNode
	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]

		if node.Pos.X == targetPos.X && node.Pos.Y == targetPos.Y {
			path := []Vector2D{}
			var pathNode = node
			for pathNode.Parent != nil {
				path = append([]Vector2D{pathNode.Pos}, path...)
				pathNode = pathNode.Parent
			}
			return path, nil
		}

		for _, validKnightMovement := range validKnightMovements {
			newPos := Vector2D{
				X: node.Pos.X + validKnightMovement.X,
				Y: node.Pos.Y + validKnightMovement.Y,
			}

			if board.isAValidPos(newPos) && !isCellVisited[newPos.X][newPos.Y] {
				isCellVisited[newPos.X][newPos.Y] = true
				newNode := &PathNode{
					Parent: node,
					Pos:    newPos,
				}
				queue = append(queue, newNode)
			}

		}
	}

	return nil, constants.ErrTargetIsUnreachable
}

func (board *Board) isAValidPos(pos Vector2D) bool {
	return pos.X >= constants.MinCellPos && pos.X < board.size && pos.Y >= constants.MinCellPos && pos.Y < board.size
}
