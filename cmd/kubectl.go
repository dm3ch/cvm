package cmd

import (
	"github.com/spf13/cobra"
)

var kubectlCmd = &cobra.Command{
	Use:   "kubectl",
	Short: "Controll kubectl version",
}

func init() {
	rootCmd.AddCommand(kubectlCmd)
}
