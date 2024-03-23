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

// updateDataCmd represents the updateData command
var updateDataCmd = &cobra.Command{
	Use:   "update-data",
	Short: "更新双色球数据信息",
	Long:  `从 https://www.cwl.gov.cn 网站爬取所有历史双色球信息,更新本地数据`,
	Run: func(cmd *cobra.Command, args []string) {
		cookie := cmd.Flag("cookie").Value.String()
		service.Update(cookie)
	},
}

func init() {
	rootCmd.AddCommand(updateDataCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	updateDataCmd.Flags().StringP("cookie", "c", "", "网站cookie 信息")
	updateDataCmd.MarkFlagRequired("cookie")
}
