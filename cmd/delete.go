package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <version 1> ... <version n>",
	Short: "Deletes kubectl version",
	Args: func(cmd *cobra.Command, args []string) error {
		isAll, err := cmd.Flags().GetBool("all")
		if err != nil {
			return err
		}

		if !isAll && len(args) < 1 {
			return fmt.Errorf("List of versions or --all option should be supplied")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var versions []string
		var err error

		exitCode := 0
		isAll, _ := cmd.Flags().GetBool("all")

		if isAll {
			versions, err = listLocal()
			if err != nil {
				fmt.Println("Error: ", err)
				os.Exit(1)
			}
		} else {
			versions = args
		}

		for _, v := range versions {
			err := deleteKubectlVersion(v)
			if err == nil {
				fmt.Printf("Successfully removed %s kubectl version.\n", v)
			} else {
				fmt.Printf("Failed to remove %s kubectl version.\n", v)
				fmt.Println("Error: ", err)
				exitCode = 1
			}
			fmt.Println()
		}

		if exitCode != 0 {
			os.Exit(exitCode)
		}
	},
}

func init() {
	kubectlCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().BoolP("all", "a", false, "Remove all local versions")
}

func deleteKubectlVersion(version string) error {
	storagePath, err := getStorageDirAbsolutePath()
	if err != nil {
		return err
	}

	err = os.Remove(filepath.Join(storagePath, "kubectl", "kubectl_"+version))
	if err != nil {
		return err
	}

	return nil
}
