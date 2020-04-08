package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultStorageDir = "~/.cvm"
	envStorageDir     = "CVM_STORAGE_DIR"
)

var rootCmd = &cobra.Command{
	Use:   "cvm",
	Short: "Single tool that allows easily install different versions cloud tools and switch beetwen them",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		storageDir, err := getStorageDirAbsolutePath()
		if err != nil {
			fmt.Println("Can't get storage directory absolute path:")
			fmt.Println(err)
			os.Exit(1)
		}

		err = createDirIfNotExist(storageDir + "/kubectl")
		if err != nil {
			fmt.Println("Can't create storage directory:")
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("storageDir", "C", defaultStorageDir,
		fmt.Sprintf("Storage directory (Could be also set via %s environment variable)", envStorageDir))
	_ = viper.BindPFlag("storageDir", rootCmd.PersistentFlags().Lookup("storageDir"))
	_ = viper.BindEnv("storageDir", envStorageDir)
	viper.SetDefault("storageDir", defaultStorageDir)
}

// Create directory if it doesn't exists
func createDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		return err
	}

	return nil
}

// Get staorage directory relative path
func getStorageDirRelativePath() string {
	return viper.GetString("storageDir")
}

// Get storage directory absolute path
func getStorageDirAbsolutePath() (string, error) {
	storageDir, err := homedir.Expand(getStorageDirRelativePath())
	if err != nil {
		return "", err
	}

	return storageDir, nil
}
