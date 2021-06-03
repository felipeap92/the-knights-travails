package board

import (
	"reflect"
	"testing"

	"github.com/felipeap92/the-knights-travails/constants"
)

func TestNewBoard(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name        string
		args        args
		expected    *Board
		expectedErr error
	}{
		{
			name:        "invalid board size",
			args:        args{-15},
			expected:    nil,
			expectedErr: constants.ErrInvalidBoardSize,
		},
		{
			name:        "invalid board size",
			args:        args{0},
			expected:    nil,
			expectedErr: constants.ErrInvalidBoardSize,
		},
		{
			name:        "invalid board size",
			args:        args{2},
			expected:    nil,
			expectedErr: constants.ErrInvalidBoardSize,
		},
		{
			name:        "create successfully",
			args:        args{3},
			expected:    &Board{3},
			expectedErr: nil,
		},
		{
			name:        "create successfully",
			args:        args{100},
			expected:    &Board{100},
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBoard(tt.args.size)
			if err != tt.expectedErr {
				t.Errorf("NewBoard() error = %v, expected %v", err, tt.expectedErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("NewBoard() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestBoard_FindShortestPathForKnight(t *testing.T) {
	type args struct {
		startPos  Vector2D
		targetPos Vector2D
	}
	tests := []struct {
		name        string
		boardSize   int
		args        args
		expected    []Vector2D
		expectedErr error
	}{
		{
			name:        "target is unreachable",
			boardSize:   constants.BoardSize,
			args:        args{Vector2D{0, 0}, Vector2D{10, 10}},
			expected:    nil,
			expectedErr: constants.ErrTargetIsUnreachable,
		},
		{
			name:        "find path successfully from A8(0,7) to B7(1,6)",
			boardSize:   constants.BoardSize,
			args:        args{Vector2D{0, 7}, Vector2D{1, 6}},
			expected:    []Vector2D{{2, 6}, {4, 7}, {3, 5}, {1, 6}},
			expectedErr: nil,
		},
		{
			name:        "find path successfully from D5(3,4) to A1(0,0)",
			boardSize:   constants.BoardSize,
			args:        args{Vector2D{3, 4}, Vector2D{0, 0}},
			expected:    []Vector2D{{4, 2}, {2, 1}, {0, 0}},
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board, _ := NewBoard(tt.boardSize)
			got, err := board.FindShortestPath(&Knight{}, tt.args.startPos, tt.args.targetPos)
			if err != tt.expectedErr {
				t.Errorf("Board.FindShortestPath() error = %v, expected %v", err, tt.expectedErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Board.FindShortestPath() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestBoard_FindShortestPathForCamel(t *testing.T) {
	type args struct {
		startPos  Vector2D
		targetPos Vector2D
	}
	tests := []struct {
		name        string
		boardSize   int
		args        args
		expected    []Vector2D
		expectedErr error
	}{
		{
			name:        "target is unreachable",
			boardSize:   constants.BoardSize,
			args:        args{Vector2D{0, 0}, Vector2D{10, 10}},
			expected:    nil,
			expectedErr: constants.ErrTargetIsUnreachable,
		},
		{
			name:        "find path successfully from A8(0,7) to B7(1,6)",
			boardSize:   constants.BoardSize,
			args:        args{Vector2D{0, 7}, Vector2D{1, 6}},
			expected:    []Vector2D{{3, 6}, {2, 3}, {1, 6}},
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board, _ := NewBoard(tt.boardSize)
			got, err := board.FindShortestPath(&Camel{}, tt.args.startPos, tt.args.targetPos)
			if err != tt.expectedErr {
				t.Errorf("Board.FindShortestPath() error = %v, expected %v", err, tt.expectedErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Board.FindShortestPath() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
