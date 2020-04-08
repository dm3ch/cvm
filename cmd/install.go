package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install <version>",
	Short: "Install kubectl version",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		version := args[0]
		url := fmt.Sprintf("https://storage.googleapis.com/kubernetes-release/release/%s/bin/%s/%s/kubectl", version, runtime.GOOS, runtime.GOARCH)
		fmt.Println("Downloading: ", url)

		r, err := http.Get(url)
		defer r.Body.Close()
		if err != nil {
			fmt.Println(err)
		}

		if r.StatusCode != http.StatusOK {
			fmt.Printf("Server returned error: %s\n", r.Status)
		}

		storagePath, err := getStorageDirAbsolutePath()
		if err != nil {
			fmt.Println("Can't get storage directory absolute path:")
			fmt.Println(err)
			os.Exit(1)
		}

		file, err := os.Create(path.Join(storagePath, "kubectl", "kubectl_"+version))
		defer file.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		_, err = io.Copy(file, r.Body)
		fmt.Printf("kubectl %s successfully installed\n", version)

		err = file.Chmod(0755)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	kubectlCmd.AddCommand(installCmd)
}
