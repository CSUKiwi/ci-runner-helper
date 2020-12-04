package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var packagePath string
var inputPath string
var atomName string
var stageIndex int
var jobIndex int
var atomIndex int
func init() {
	atomBeforeCmd.Flags().StringVarP(&atomName,"atomName","a","","atomName")
	atomBeforeCmd.Flags().StringVarP(&packagePath,"packagePath","b","","packagePath")
	atomBeforeCmd.Flags().StringVarP(&inputPath,"inputPath","c","","inputPath")
	atomBeforeCmd.Flags().IntVarP(&stageIndex, "stageIndex", "d", 0, "stageIndex")
	atomBeforeCmd.Flags().IntVarP(&jobIndex, "jobIndex", "e", 0, "jobIndex")
	atomBeforeCmd.Flags().IntVarP(&atomIndex, "atomIndex", "f", 0, "atomIndex")

	rootCmd.AddCommand(atomBeforeCmd)
}

var atomBeforeCmd = &cobra.Command{
	Use:   "atom_before",
	Short: "atom before",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("i am atom before")
		var err error
		CI_WORKSPACE := os.Getenv("CI_WORKSPACE")
		// 根据atom目录规则
		ci_data_dir := fmt.Sprintf("%s/stage-%d/job-%d/atom-%d",
			CI_WORKSPACE,stageIndex,jobIndex,atomIndex)

		if Exists(ci_data_dir) == false {
			// 创建目录
			err = os.MkdirAll(ci_data_dir, os.ModePerm)
			if err != nil {
				fmt.Println("Error creating directory")
				os.Exit(1)
			}
		}

		// 下载原子
		savefile := fmt.Sprintf("%s/%s",ci_data_dir,atomName)
		if Exists(savefile) == false {
			err = Download(packagePath,savefile)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Download url:%s savefile: %s error:%v\n ",
					packagePath,savefile,err)
				os.Exit(1)
			}
		}

		// 下载原子 input json
		savefile = fmt.Sprintf("%s/%s",ci_data_dir,"input.json")
		if Exists(savefile) == false {
			err = Download(inputPath,savefile)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Download url:%s savefile: %s error:%v\n ",
					packagePath,savefile,err)
				os.Exit(1)
			}
		}
	},
}

