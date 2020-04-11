package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
)

const releasesURL string = "https://api.github.com/repos/kubernetes/kubernetes/releases?per_page=100"
const perPage int = 100

type kubeVersion struct {
	Name string `json:"name"`
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available",
	Run: func(cmd *cobra.Command, args []string) {
		var versions []string

		isRemote, err := cmd.Flags().GetBool("remote")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		if isRemote {
			versions, err = listRemote()
		} else {
			versions, err = listLocal()
		}

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
	listCmd.Flags().BoolP("remote", "r", false, "List remote versions")
}

func listRemote() ([]string, error) {
	r, err := http.Get(releasesURL)
	defer r.Body.Close()

	if err != nil {
		return nil, err
	}

	var kubeVersions []kubeVersion
	err = json.NewDecoder(r.Body).Decode(&kubeVersions)
	if err != nil {
		return nil, err
	}

	var versions []string
	for _, v := range kubeVersions {
		versions = append(versions, v.Name)
	}
	return versions, nil
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
