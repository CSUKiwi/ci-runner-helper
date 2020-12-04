package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(artifactCmd)
}

var artifactCmd = &cobra.Command{
	Use:   "artifact",
	Short: "artifact",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("i am artifact cmd")
	},
}
