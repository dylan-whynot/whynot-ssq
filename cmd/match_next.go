/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"

	"github.com/dylan_whynot/whynot_ssq/service"

	"github.com/spf13/cobra"
)

// matchNextCmd represents the matchNext command
var matchNextCmd = &cobra.Command{
	Use:   "match-next",
	Short: "根据输入的篮球号码 统计下一期每种球出现的次数",
	Long:  `根据输入的蓝色球号码 统计下一期每一种球出现的次数.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(inputBlue) == 0 {
			log.Fatalln("blue is required")
		}
		result := service.MatchNext(query, condition)
		service.PrintMatchNextResult(query, condition, printControl, result)
	},
}

func init() {
	rootCmd.AddCommand(matchNextCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// matchNextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// matchNextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
