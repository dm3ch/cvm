package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install <version 1> ... <version n>",
	Short: "Install kubectl version",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		exitCode := 0

		for _, v := range args {
			fmt.Printf("Downloading %s kubectl version.\n", v)

			err := downloadKubectlVersion(v)
			if err == nil {
				fmt.Printf("Successfully installed %s kubectl version.\n", v)
			} else {
				fmt.Printf("Failed to install %s kubectl version.\n", v)
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
	kubectlCmd.AddCommand(installCmd)
}

func downloadKubectlVersion(version string) error {
	url := fmt.Sprintf("https://storage.googleapis.com/kubernetes-release/release/%s/bin/%s/%s/kubectl", version, runtime.GOOS, runtime.GOARCH)

	r, err := http.Get(url)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("Server returned error: %s\n", r.Status)
	}

	storagePath, err := getStorageDirAbsolutePath()
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(storagePath, "kubectl", "kubectl_"+version))
	defer file.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(file, r.Body)
	if err != nil {
		return err
	}

	err = file.Chmod(0755)
	if err != nil {
		return err
	}

	return nil
}
