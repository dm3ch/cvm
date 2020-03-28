package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "cvm",
	Short: "Single tool that allows easily install different versions cloud tools and switch beetwen them",
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
