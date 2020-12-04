package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var pipeline string
var outputPath string


type AtomJson struct {
	Pipeline   string     `json:"pepeline"`
	StageIndex int        `json:"stage_index"`
	JobIndex   int        `json:"job_index"`
	AtomIndex  int        `json:"atom_index"`
	AtomOutput AtomOutput `json:"atom_output"`
}
type AtomOutput struct {
	Status    string                 `json:"status"`
	Message   string                 `json:"message"`
	ErrorCode int                    `json:"errorCode"`
	Type      string                 `json:"type"`
	Data      map[string]interface{} `json:"data"`
}

func init() {

	atomAfterCmd.Flags().IntVarP(&stageIndex, "stageIndex", "d", 0, "stageIndex")
	atomAfterCmd.Flags().IntVarP(&jobIndex, "jobIndex", "e", 0, "jobIndex")
	atomAfterCmd.Flags().IntVarP(&atomIndex, "atomIndex", "f", 0, "atomIndex")

	atomAfterCmd.Flags().StringVarP(&pipeline, "pipeline", "g", "", "pipeline")
	atomAfterCmd.Flags().StringVarP(&outputPath,"outputPath","i","","outputPath")


	rootCmd.AddCommand(atomAfterCmd)
}

var atomAfterCmd = &cobra.Command{
	Use:   "atom_after",
	Short: "atom after",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("i am atom after")
		var err error
		CI_WORKSPACE := os.Getenv("CI_WORKSPACE")
		// 根据atom目录规则
		output_file := fmt.Sprintf("%s/stage-%d/job-%d/atom-%d/output.json",
			CI_WORKSPACE, stageIndex, jobIndex, atomIndex)

		atomJson := &AtomJson{
			Pipeline:   pipeline,
			StageIndex: stageIndex,
			JobIndex:   jobIndex,
			AtomIndex:  atomIndex,
		}

		var atomOutput AtomOutput
		data, err := ioutil.ReadFile(output_file)
		if err != nil {
			fmt.Println("read output err:",err)
			os.Exit(1)
		}
		err = json.Unmarshal(data, &atomOutput)
		if err != nil {
			fmt.Println("json Unmarshal err:",err)
			os.Exit(1)
		}
		atomJson.AtomOutput = atomOutput

		json, err := json.Marshal(atomJson)
		if err != nil {
			fmt.Println("生成json字符串错误")
			os.Exit(1)

		}
		err = PostJson(outputPath,json)
		if err != nil {
			fmt.Println("PostJson faild")
			os.Exit(1)
		}
	},
}
