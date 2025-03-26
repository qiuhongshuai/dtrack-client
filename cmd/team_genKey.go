/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"dtrack-client/types"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// genKeyCmd represents the genKey command
var teamGenKeyCmd = &cobra.Command{
	Use:   "genKey",
	Short: "Generate an API key for a specified team",
	Long: `Generate an API key for a specific team by providing its UUID. If no UUID is provided, the command will generate an API key for the currently authenticated team.

- Use the --id flag to specify the UUID of the team for which you want to generate the API key. This flag is optional.
- If the --id flag is not provided, the command will automatically use the UUID of the team associated with the current authentication.

Example usage:
- To generate an API key for a specific team with UUID "12345678-1234-1234-1234-1234567890ab":  
  dtrack team genKey --id 12345678-1234-1234-1234-1234567890ab
- To generate an API key for the currently authenticated team:  
  dtrack team genKey

Note: Ensure that you have the necessary permissions to generate API keys for the specified team. If the generation fails due to permission issues or other errors, an appropriate error message will be displayed.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("genKey called")
		id := cmd.Flag("id").Value.String()
		if strings.TrimSpace(id) == "" {
			fmt.Println("Warning:未提供id，使用当前认证apikey的team uuid")
			//return
			var team types.Team
			res, err := Client.R().Execute(types.ApiUrls[types.SelfTeam].Method, types.ApiUrls[types.SelfTeam].Url)
			if err != nil {
				fmt.Printf("获取当前团队报错：%s", err.Error())
				return
			}
			if !res.IsSuccess() {
				fmt.Printf("获取当前团队失败：%s", res.String())
				return
			}
			if err := json.Unmarshal(res.Bytes(), &team); err != nil {
				fmt.Printf("获取当前团队报错：%s", err.Error())
				return
			}
			id = team.Uuid
		}
		defer Client.Close()
		res, err := Client.R().SetPathParam("id", id).Execute(types.ApiUrls[types.GenTeamKey].Method, types.ApiUrls[types.GenTeamKey].Url)
		if err != nil {
			fmt.Printf("生成团队apikey报错：%s", err.Error())
			return
		}
		if !res.IsSuccess() {
			fmt.Printf("生成团队apikey失败：%s", res.Status())
			return
		}
		fmt.Println(res.String())
	},
}

func init() {
	teamCmd.AddCommand(teamGenKeyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genKeyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genKeyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	teamGenKeyCmd.Flags().String("id", "", "team uuid")

}
