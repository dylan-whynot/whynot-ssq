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
package main

import "github.com/dylan_whynot/whynot_ssq/cmd"

func main() {
	cmd.Execute()
	//db.LoadDatas()
	//query := &model.Query{Issues: []string{"2024026", "2013030"}, Blue: []string{}, Week: "", StartYear: "", EndYear: ""}
	//printControl := &model.PrintControl{PageSize: 50, GranterThan: 0, PrintIssues: false}
	////condition := &model.Condition{BallColor: "blue", RedCount: 1, RangeSize: 5}
	//result := service.Search(query)
	//service.PrintSearchResult(query, printControl, result)
}
