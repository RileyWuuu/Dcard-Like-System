package cmd

import (
	"dcard/service/apigateway"

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
	apigateway.EnableApiGateway()
	return nil
}
