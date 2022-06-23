package cmd

import (
	"dcard/service/post"

	"github.com/spf13/cobra"
)

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "none.",
	Long:  `No more description.`,
	RunE:  RunPostCmd,
}

func init() {
	rootCmd.AddCommand(postCmd)
}

func RunPostCmd(cmd *cobra.Command, args []string) error {
	post.EnablePostServer()
	return nil
}
