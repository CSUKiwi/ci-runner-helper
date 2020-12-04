package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(gitCmd)
}

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "git clone source code",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("i am git cmd")

		CI_PROJECT_URL := os.Getenv("CI_PROJECT_URL")
		CI_PROJECT_BRANCH := os.Getenv("CI_PROJECT_BRANCH")
		CI_PROJECT_DIR := os.Getenv("CI_PROJECT_DIR")

		if Exists(CI_PROJECT_DIR) {
			fmt.Println("CI_PROJECT_DIR Exists.")
			os.Exit(0)
		}

		//CI_PROJECT_URL = "https://gitee.com/bruce_luotao/maven-demo.git"
		//CI_PROJECT_BRANCH = "main"
		//CI_PROJECT_DIR = "/tmp/xxyyzz/fdev-ci/maven-demo"

		script := fmt.Sprintf("git clone %s -b %s %s",
			CI_PROJECT_URL,
			CI_PROJECT_BRANCH,
			CI_PROJECT_DIR)



		fmt.Println(script)
		output, err := ExecuteCommand("git", "clone",CI_PROJECT_URL ,"-b",CI_PROJECT_BRANCH,CI_PROJECT_DIR)
		if err != nil {
				fmt.Fprintf(os.Stderr, "execute %s  error:%v output:%s\n ", script, err, output)
				os.Exit(1)
		}
		fmt.Println(output)

	},
}
