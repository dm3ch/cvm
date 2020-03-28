package cmd

import (
	"os"
	"fmt"
	"net/http"
	"encoding/json"

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
		r, err := http.Get(releasesURL)


		if err != nil {
			fmt.Printf(err.Error())
			os.Exit(1)
		}
		defer r.Body.Close()
		
		var versions []kubeVersion
		err = json.NewDecoder(r.Body).Decode(&versions)
		if err != nil {
			fmt.Printf(err.Error())
			os.Exit(1)
		}
		
		for _, v := range versions {
			fmt.Println(v.Name)
		}
	},
}

func init() {
	kubectlCmd.AddCommand(listCmd)
}
