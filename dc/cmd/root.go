package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/hawkv6/dummy-controller/internal/config"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "dc",
	Short: "dc is a dummy-controller for testing the hawkwing agent.",
	Long: `dc is a dummy-controller for testing the hawkwing agent.
	Complete documentation is available at:
	https://github.com/hawkv6/hawkwing
	`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		config.GetInstance().SetConfigFile(cfgFile)
	}

	if err := config.Parse(); err != nil {
		log.Fatalln(err)
	}
}
