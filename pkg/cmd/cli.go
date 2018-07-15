package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd root command
var rootCmd = &cobra.Command{
	Use:   "gosdf",
	Short: "SDF Yaml Utility",
	Long:  ``,
}

// Execute executes rootCmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
