package cmd

import (
	"dcard/matching"

	"github.com/spf13/cobra"
)

var matchingCmd = &cobra.Command{
	Use:   "matching",
	Short: "none.",
	Long:  `No more description.`,
	RunE:  RunMatchingCmd,
}

func init() {
	rootCmd.AddCommand(matchingCmd)
}

func RunMatchingCmd(cmd *cobra.Command, args []string) error {
	matching.EnableMatchingServer()
	return nil
}
