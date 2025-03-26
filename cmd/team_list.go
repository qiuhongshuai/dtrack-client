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
var teamListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all teams with optional formatting",
	Long: `Retrieve and display a list of all teams in the system. The output can be formatted as either JSON or a table for easier readability.

- Use the --format flag to specify the output format. Supported values are "json" (default) and "table".
- If no format is specified, the command will default to JSON output.

Example usage:
- To list all teams in JSON format:  
  dtrack team list
- To list all teams in table format:  
  dtrack team list --format table

Note: If the command encounters an error while fetching or parsing the team data, it will display an appropriate error message. Ensure that you have the necessary permissions to view the team list.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("list called")
		format := cmd.Flag("format").Value.String()

		defer Client.Close()
		res, err := Client.R().Execute(types.ApiUrls[types.ListTeams].Method, types.ApiUrls[types.ListTeams].Url)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		if !res.IsSuccess() {
			fmt.Printf("Error: %s\n", res.String())
		}
		if format == "json" {
			fmt.Println(res.String())
		} else {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"uuid", "name"})
			var teams []*types.Team
			if err := json.Unmarshal(res.Bytes(), &teams); err != nil {
				fmt.Printf("Error: %s\n", err)
			}
			for _, team := range teams {
				table.Append([]string{team.Uuid, team.Name})
			}
			table.Render()
		}
	},
}

func init() {
	teamCmd.AddCommand(teamListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	teamListCmd.Flags().String("format", "json", "输出格式，支持table，json")
	teamListCmd.Flag("format").Annotations = map[string][]string{
		cobra.BashCompCustom: {"__cobra_comp_format"},
	}
	teamListCmd.RegisterFlagCompletionFunc("format", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"table", "json"}, cobra.ShellCompDirectiveDefault
	})
}
