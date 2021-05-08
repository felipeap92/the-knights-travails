package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/felipeap92/the-knights-travails/board"
)

func CheckIfStringIsAValidAlgebraicChessNotation(chessNotation string) bool {
	matched, _ := regexp.MatchString("^[A-H][1-8]$", chessNotation)
	return matched
}

func ConvertAlgebraicChessNotationToVector2D(chessNotation string) board.Vector2D {
	return board.Vector2D{
		X: int(chessNotation[0] - 'A'),
		Y: int(chessNotation[1]-'0') - 1,
	}
}

func ConvertVector2DToAlgebraicChessNotation(vector board.Vector2D) string {
	letter := string(rune(vector.X + 'A'))
	number := string(rune(vector.Y+'0') + 1)
	return fmt.Sprint(letter, number)
}

func ConvertVector2DSliceToAlgebraicChessNotationSlice(slice []board.Vector2D) []string {
	sliceInChessNotation := make([]string, len(slice))
	for index, pos := range slice {
		sliceInChessNotation[index] = ConvertVector2DToAlgebraicChessNotation(pos)
	}
	return sliceInChessNotation
}

func PrintPathInAlgebraicChessNotation(path []board.Vector2D) {
	pathInAlgebraicChessNotation := ConvertVector2DSliceToAlgebraicChessNotationSlice(path)
	fmt.Println(strings.Join(pathInAlgebraicChessNotation, " "))
}
