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

// searchCmd represents the search command
var matchCountCmd = &cobra.Command{
	Use:   "match-count",
	Short: "提供统计功能",
	Long: `根据输入参数
    eg: -b 5 -s 2020 -e 2024 -r 3 表示在[2020,2024)年内 蓝球号码为5时,统计3个红球组合出现的频率`,
	Run: func(cmd *cobra.Command, args []string) {
		if redCount == 0 {
			log.Fatalln("red-count granter than 0")
		}
		times := service.MatchCount(query, condition)
		service.PrintMatchCountResult(query, condition, printControl, times)
	},
}

func init() {
	rootCmd.AddCommand(matchCountCmd)
}
