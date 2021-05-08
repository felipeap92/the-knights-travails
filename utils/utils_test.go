package utils

import (
	"reflect"
	"testing"

	"github.com/felipeap92/the-knights-travails/board"
)

func TestCheckIfStringIsAValidAlgebraicChessNotation(t *testing.T) {
	tests := []struct {
		name          string
		chessNotation string
		isValid       bool
	}{
		{
			name:          "invalid chess notation",
			chessNotation: "A10",
			isValid:       false,
		},
		{
			name:          "invalid chess notation",
			chessNotation: "J1",
			isValid:       false,
		},
		{
			name:          "invalid chess notation",
			chessNotation: "B01",
			isValid:       false,
		},
		{
			name:          "valid chess notation",
			chessNotation: "A1",
			isValid:       true,
		},
		{
			name:          "valid chess notation",
			chessNotation: "B8",
			isValid:       true,
		},
		{
			name:          "valid chess notation",
			chessNotation: "H7",
			isValid:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckIfStringIsAValidAlgebraicChessNotation(tt.chessNotation); got != tt.isValid {
				t.Errorf("CheckIfStringIsAValidAlgebraicChessNotation(%v) = %v, expected %v", tt.chessNotation, got, tt.isValid)
			}
		})
	}
}

func TestConvertAlgebraicChessNotationToVector2D(t *testing.T) {
	tests := []struct {
		name          string
		chessNotation string
		expected      board.Vector2D
	}{
		{
			name:          "convert A1",
			chessNotation: "A1",
			expected:      board.Vector2D{0, 0},
		},
		{
			name:          "convert A1",
			chessNotation: "H8",
			expected:      board.Vector2D{7, 7},
		},
		{
			name:          "convert A1",
			chessNotation: "C5",
			expected:      board.Vector2D{2, 4},
		},
		{
			name:          "convert A1",
			chessNotation: "F2",
			expected:      board.Vector2D{5, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertAlgebraicChessNotationToVector2D(tt.chessNotation); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ConvertAlgebraicChessNotationToVector2D() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestConvertVector2DToAlgebraicChessNotation(t *testing.T) {
	tests := []struct {
		name     string
		vector   board.Vector2D
		expected string
	}{
		{
			name:     "convert A1",
			vector:   board.Vector2D{0, 0},
			expected: "A1",
		},
		{
			name:     "convert H8",
			vector:   board.Vector2D{7, 7},
			expected: "H8",
		},
		{
			name:     "convert C5",
			vector:   board.Vector2D{2, 4},
			expected: "C5",
		},
		{
			name:     "convert F2",
			vector:   board.Vector2D{5, 1},
			expected: "F2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertVector2DToAlgebraicChessNotation(tt.vector); got != tt.expected {
				t.Errorf("ConvertVector2DToAlgebraicChessNotation() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestConvertVector2DSliceToAlgebraicChessNotationSlice(t *testing.T) {
	tests := []struct {
		name     string
		slice    []board.Vector2D
		expected []string
	}{
		{
			name:     "convert successfully",
			slice:    []board.Vector2D{{0, 0}, {7, 7}, {2, 4}, {5, 1}},
			expected: []string{"A1", "H8", "C5", "F2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertVector2DSliceToAlgebraicChessNotationSlice(tt.slice); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ConvertVector2DSliceToAlgebraicChessNotationSlice() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
