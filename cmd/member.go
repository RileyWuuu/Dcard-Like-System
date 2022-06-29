package cmd

import (
	"dcard/service/member"
	"dcard/storage/mongo"
	"dcard/storage/mysql"
	"dcard/storage/redis"

	"github.com/spf13/cobra"
)

var memberCmd = &cobra.Command{
	Use:   "member",
	Short: "none.",
	Long:  `No more description.`,
	RunE:  RunMemberCmd,
}

func init() {
	rootCmd.AddCommand(memberCmd)
}

func RunMemberCmd(cmd *cobra.Command, args []string) error {
	mysql.Initialize()
	redis.Initialize()
	mongo.Initialize()

	member.EnableMemberServer()

	return nil
}
