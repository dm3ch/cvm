package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <version 1>...<version n>",
	Short: "Deletes kubectl version",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		version := args[0]

		storagePath, err := getStorageDirAbsolutePath()
		if err != nil {
			fmt.Println("Can't get storage directory absolute path:")
			fmt.Println(err)
			os.Exit(1)
		}

		err = os.Remove(filepath.Join(storagePath, "kubectl", "kubectl_"+version))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Kubectl versions removed")
	},
}

func init() {
	kubectlCmd.AddCommand(deleteCmd)
}
