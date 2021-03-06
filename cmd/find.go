package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/felipeap92/the-knights-travails/board"
	"github.com/felipeap92/the-knights-travails/constants"
	"github.com/felipeap92/the-knights-travails/utils"
	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:     "find",
	Short:   "Find the shortest sequence of moves for a Knight to move between the starting and the target positions.",
	Example: fmt.Sprint(rootCmd.Name(), " find A8 B7"),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 3 {
			return errors.New("find requires 3 arguments, the starting and target position in algebraic chess notation")
		}

		if !utils.CheckIfStringIsAValidAlgebraicChessNotation(args[1]) || !utils.CheckIfStringIsAValidAlgebraicChessNotation(args[2]) {
			return errors.New("the starting and target position should be in algebraic chess notation")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		pieceName := args[0]
		startPos := utils.ConvertAlgebraicChessNotationToVector2D(args[1])
		targetPost := utils.ConvertAlgebraicChessNotationToVector2D(args[2])

		piece, err := board.CreateNewPiece(pieceName)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		board, err := board.NewBoard(constants.BoardSize)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		path, err := board.FindShortestPath(piece, startPos, targetPost)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		utils.PrintPathInAlgebraicChessNotation(path)
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(findCmd)
}
