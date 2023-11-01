package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hawkv6/dummy-controller/pkg/messaging"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start dummy-controller server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting dummy-controller server...")
		server := messaging.NewMessagingServer()
		server.Start()
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c

		fmt.Println("\nDummy-controller stopped")
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
