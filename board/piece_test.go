package board

import (
	"reflect"
	"testing"

	"github.com/felipeap92/the-knights-travails/constants"
)

func TestCreateNewPiece(t *testing.T) {
	type args struct {
		pieceName string
	}
	tests := []struct {
		name        string
		args        args
		expected    Piece
		expectedErr error
	}{
		{
			name:        "piece does not exist",
			args:        args{"invalid-piece"},
			expected:    nil,
			expectedErr: constants.ErrPieceDoesNotExist,
		},
		{
			name:        "knight piece",
			args:        args{"knight"},
			expected:    &Knight{},
			expectedErr: nil,
		},
		{
			name:        "camel piece",
			args:        args{"camel"},
			expected:    &Camel{},
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateNewPiece(tt.args.pieceName)
			if err != tt.expectedErr {
				t.Errorf("CreateNewPiece() error = %v, expected %v", err, tt.expectedErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("CreateNewPiece() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
