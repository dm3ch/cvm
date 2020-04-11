package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
)

type kubeVersion struct {
	Name string `json:"name"`
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List versions",
	Long:  "Lists local versions by default.\nUse \"list remote\" to list all available versions",
	Run: func(cmd *cobra.Command, args []string) {
		var versions []string

		versions, err := listLocal()

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		for _, v := range versions {
			fmt.Println(v)
		}
	},
}

func init() {
	kubectlCmd.AddCommand(listCmd)
}

func listLocal() ([]string, error) {
	storagePath, err := getStorageDirAbsolutePath()
	if err != nil {
		return nil, err
	}

	files, err := filepath.Glob(filepath.Join(storagePath, "kubectl", "kubectl_*"))

	var versions []string
	r := regexp.MustCompile(`kubectl_(v\d+.\d+.\d+)$`)
	for _, f := range files {
		submatches := r.FindStringSubmatch(f)
		if len(submatches) == 2 {
			versions = append(versions, submatches[1])
		}
	}
	return versions, nil
}
