package cmd

import (
	"os"

	"github.com/pingkuan/backendhw/server"
	"github.com/spf13/cobra"
)

var port int

var rootCmd = &cobra.Command{
	Use:   "backendhw",
	Short: "Storing line messages to mongoDB",
	Long:  `An api receiving messages from line platform.`,
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start gin server",
	Long:  `start gin server with given port.`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Server(port)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().IntVar(&port, "port", 8080, "Default port on 8080")
}
