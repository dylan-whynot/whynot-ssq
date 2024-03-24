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
	"github.com/dylan_whynot/whynot_ssq/service"
	"github.com/spf13/cobra"
)

var issues []string

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "提供对数据的检索和基本的查询功能",
	Long:  `提供对数据的检索和基本的查询功能.`,
	Run: func(cmd *cobra.Command, args []string) {
		result := service.Search(query)
		service.PrintSearchResult(query, printControl, result)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	searchCmd.Flags().StringArrayVar(&issues, "issues", []string{}, "需要搜索的期号")
}
