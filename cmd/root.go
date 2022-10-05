package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Addr string
var Port string
var VERSION string = "0.0.1"

var rootCmd = &cobra.Command{
	Use:   "simple-kv",
	Short: "smiple-kv is a simple key-value service",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(VERSION)
		os.Exit(0)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&Addr, "address", "a", "127.0.0.1", "the listen address, default is 127.0.0.1")
	rootCmd.PersistentFlags().StringVarP(&Port, "port", "p", "8224", "the listen port, default is 8224")
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
