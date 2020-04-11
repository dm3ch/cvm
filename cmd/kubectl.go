package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var kubectlCmd = &cobra.Command{
	Use:   "kubectl",
	Short: "Controll kubectl version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello!")
	},
}

func init() {
	rootCmd.AddCommand(kubectlCmd)
}
