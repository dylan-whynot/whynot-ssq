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
	"fmt"
	"os"

	"github.com/dylan_whynot/whynot_ssq/service"

	"github.com/spf13/cobra"
)

var rangeSize int

// matchRangeCmd represents the matchRange command
var matchRangeCmd = &cobra.Command{
	Use:   "match-range",
	Short: "统计当蓝球为 x时红球落在某个区间的数量",
	Long:  `统计当蓝球为 x时红球落在某个区间的数量`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if rangeSize < 2 || rangeSize > 33 {
			fmt.Println("range-size [2,33]")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		condition.RangeSize = rangeSize
		result := service.MatchRange(query, condition)
		service.PrintMatchRangeResult(query, condition, printControl, result)
	},
}

func init() {
	rootCmd.AddCommand(matchRangeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// matchRangeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	matchRangeCmd.Flags().IntVar(&rangeSize, "range-size", 5, "设置区间范围 最小值为2 最大值为33")
}
