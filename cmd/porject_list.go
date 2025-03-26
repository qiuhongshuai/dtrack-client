/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"dtrack-client/types"
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available projects with optional filtering and formatting",
	Long:  `List available projects with optional filtering and formatting`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("list called")
		name := cmd.Parent().Flag("name").Value.String()
		simple, err := cmd.Flags().GetBool("simple")
		if err != nil {
			fmt.Printf("simple参数解析失败：%s,使用默认值", err.Error())
			simple = true
		}
		format, err := cmd.Flags().GetString("format")
		if err != nil {
			fmt.Printf("format参数解析失败：%s,使用默认值", err.Error())
			format = "json"
		}
		maxLines, err := cmd.Flags().GetInt("max-lines")
		if err != nil {
			fmt.Printf("max-lines参数解析失败：%s,使用默认值", err.Error())
			maxLines = 5
		}
		if maxLines > 100 {
			maxLines = 100
		}
		defer Client.Close()
		r := Client.R()
		if name != "" {
			r.SetQueryParam("name", name)
		}
		res, err := r.SetQueryParams(map[string]string{"pageNumber": "1", "pageSize": fmt.Sprintf("%d", maxLines)}).Execute(types.ApiUrls[types.ListProjects].Method, types.ApiUrls[types.ListProjects].Url)
		if err != nil || !res.IsSuccess() {
			fmt.Println(err.Error(), res.Status())
			return
		}
		var rr interface{}
		if err := json.Unmarshal(res.Bytes(), &rr); err != nil {
			fmt.Println(err.Error())
			return
		}
		//fmt.Printf("%T,%+v", rr, rr)
		if simple {
			if _, ok := rr.([]interface{}); ok {
				table := tablewriter.NewWriter(os.Stdout)
				if format == "table" {
					table.SetHeader([]string{"name", "uuid", "version", "classifier"})
					//table.SetAutoMergeCells(true)
					table.SetRowLine(true)
				} else {
					fmt.Println("[")
				}
				for _, v := range rr.([]interface{}) {
					result := make(map[string]interface{})
					if v1, ok := v.(map[string]interface{}); ok {
						result["name"] = v1["name"]
						result["uuid"] = v1["uuid"]
						result["version"] = v1["version"]
						result["classifier"] = v1["classifier"]
					}
					if format == "table" {
						if result["version"] == nil {
							result["version"] = ""
						}
						table.Append([]string{result["name"].(string), result["uuid"].(string), result["version"].(string), result["classifier"].(string)})
					} else {
						data, _ := json.Marshal(result)
						fmt.Println(string(data))
					}
				}
				if format == "table" {
					table.Render()
				} else {
					fmt.Println("]")
				}
			} else {
				//fmt.Println(res.String())
			}
		} else {
			fmt.Println(res.String())
		}
	},
}

func init() {
	projectCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")
	listCmd.Flags().Bool("simple", true, "是否简化输出，只显示项目名，ID，版本、团队、分类器")
	listCmd.Flags().String("format", "json", "输出格式，支持table，json，仅在simple为true时生效")
	listCmd.Flag("format").Annotations = map[string][]string{
		cobra.BashCompCustom: {"__cobra_comp_format"},
	}
	listCmd.Flags().IntP("max-lines", "l", 5, "最多输出行数，默认是5个，最大是100个")
	listCmd.RegisterFlagCompletionFunc("format", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"table", "json"}, cobra.ShellCompDirectiveDefault
	})

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
