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

	"github.com/dylan_whynot/whynot_ssq/model"

	"github.com/dylan_whynot/whynot_ssq/db"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string
var inputBlue []string
var startYear string
var endYear string
var week string
var pageSize int
var printIssues bool
var redCount int
var granterThan int
var ballColor string

var query *model.Query
var printControl *model.PrintControl
var condition *model.Condition

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "whynot-ssq",
	Short: "双色球历史数据自定义搜索功能",
	Long:  `基于历年双色球历史数据,提供查询和聚合搜索功能.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		query = &model.Query{Blue: inputBlue, Week: week, StartYear: startYear, EndYear: endYear}
		printControl = &model.PrintControl{PageSize: pageSize, GranterThan: granterThan, PrintIssues: printIssues}
		condition = &model.Condition{BallColor: ballColor, RedCount: redCount}
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	db.LoadDatas()
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.whynot_ssq.yaml)")

	rootCmd.PersistentFlags().StringArrayVarP(&inputBlue, "blue", "b", []string{}, "需要统计或者过滤的蓝球号码")
	rootCmd.PersistentFlags().StringVarP(&startYear, "start-year", "s", "", "统计开始年份包含")
	rootCmd.PersistentFlags().StringVarP(&endYear, "end-year", "e", "", "统计结束年份不包含")
	rootCmd.PersistentFlags().StringVarP(&week, "week", "w", "", "星期")
	rootCmd.PersistentFlags().IntVarP(&pageSize, "pagesize", "p", 50, "打印多少条数据 -1表示不限制")
	rootCmd.PersistentFlags().BoolVarP(&printIssues, "print-issues", "i", false, "打印期号")
	rootCmd.PersistentFlags().IntVarP(&redCount, "red-count", "r", 1, "几个红球组合出现")
	rootCmd.PersistentFlags().IntVarP(&granterThan, "granter-than", "g", 0, "输出结果时筛选出现次数大于n")
	rootCmd.PersistentFlags().StringVar(&ballColor, "ball-color", "blue", "球的颜色")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".whynot_ssq" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".whynot_ssq")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
