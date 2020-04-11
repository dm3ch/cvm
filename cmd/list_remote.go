package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

const releasesURL string = "https://api.github.com/repos/kubernetes/kubernetes/releases?per_page=100"
const perPage int = 100

var listRemoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "List available versions",
	Run: func(cmd *cobra.Command, args []string) {
		var versions []string

		versions, err := listRemote()

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
	listCmd.AddCommand(listRemoteCmd)
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
