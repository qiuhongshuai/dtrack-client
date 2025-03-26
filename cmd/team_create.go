/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"dtrack-client/types"
	"fmt"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var teamCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new team with the specified name",
	Long: `Create a new team by providing a unique name. This command allows you to add new teams to the system, which can then be used to manage access and permissions for projects.

- The team name must be provided as a single argument when executing the command.
- If no name is provided or if multiple arguments are passed, the command will return an error.

Example usage:
- To create a team named "Development":  
     dtrack team create Development

Note: Ensure that the team name is unique within the system. If the creation fails due to conflicts or other issues, an appropriate error message will be displayed.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cobra.CheckErr("请输入团队名称")
		} else if len(args) > 1 {
			cobra.CheckErr("参数过多")
		}
		teamName := args[0]
		defer Client.Close()
		res, err := Client.R().SetBody(map[string]string{"name": teamName}).Execute(types.ApiUrls[types.CreateTeam].Method, types.ApiUrls[types.CreateTeam].Url)
		if err != nil {
			fmt.Printf("创建团队失败:%s\n", err)
			return
		}
		if !res.IsSuccess() {
			fmt.Printf("创建团队失败:%s\n", res.Status())
			return
		}
		fmt.Printf("创建团队成功:%s\n", res.String())
	},
}

func init() {
	teamCmd.AddCommand(teamCreateCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//teamCreateCmd.Flags().String("parent-id", "", "父项目ID，可以为空")
}
