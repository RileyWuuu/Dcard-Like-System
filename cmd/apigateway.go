package cmd

import (
	"dcard/service/apigateway"
	"dcard/storage/mongo"
	"dcard/storage/mysql"
	"dcard/storage/redis"

	"github.com/spf13/cobra"
)

var apiGatewayCmd = &cobra.Command{
	Use:   "apigateway",
	Short: "none.",
	Long:  `No more description.`,
	RunE:  RunApiGatewayCmd,
}

func init() {
	rootCmd.AddCommand(apiGatewayCmd)
}

func RunApiGatewayCmd(cmd *cobra.Command, args []string) error {
	mysql.Initialize()
	redis.Initialize()
	mongo.Initialize()

	apigateway.EnableApiGateway()

	return nil
}
