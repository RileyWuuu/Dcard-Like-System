package cmd

import (
	"dcard/config"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "f", "./conf.d/env.yaml", "config file")
	rootCmd.PersistentFlags().StringVarP(&buildTime, "buildTime", "b", time.Now().String(), "binary build time")
}

func initConfig() {
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("Dcard")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("ReadInConfig file failed: %v\n", err)
	} else {
		fmt.Printf("Used config file: %v\n", viper.ConfigFileUsed())
	}
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

	c, err := config.NewFromViper()
	if err != nil {
		fmt.Printf("Initialize config failed: %v\n", err)
	} else {
		config.SetConfig(c)
	}

	return nil
}
