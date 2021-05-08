package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "the-knights-travails",
	Short: "This program solves The Knight’s Travails problem.",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
