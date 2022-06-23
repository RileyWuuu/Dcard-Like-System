package cmd

import (
	"dcard/service/matching"
	"dcard/storage/mongo"
	"dcard/storage/mysql"
	"dcard/storage/redis"

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
	mysql.Initialize()
	redis.Initialize()
	mongo.Initialize()

	matching.EnableMatchingServer()

	return nil
}
