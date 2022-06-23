package cmd

import (
	"fmt"
	"runtime"
	"time"

	"github.com/spf13/cobra"
)

var cfgFile string
var buildTime string

var rootCmd = &cobra.Command{
	Use:               "root",
	Short:             "description.",
	Long:              `this is an example for Cobra framework.`,
	PersistentPreRunE: PersistentPreRunBeforeCommandStartUp,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&buildTime, "buildTime", "b", time.Now().String(), "binary build time")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func PersistentPreRunBeforeCommandStartUp(cmd *cobra.Command, args []string) error {
	goVersion := runtime.Version()
	osName := runtime.GOOS
	architecture := runtime.GOARCH
	fmt.Println("======")
	fmt.Printf("Build on %s\n", buildTime)
	fmt.Printf("GoVersion: %s\n", goVersion)
	fmt.Printf("OS: %s\n", osName)
	fmt.Printf("Architecture: %s\n", architecture)
	fmt.Println("======")

	return nil
}
