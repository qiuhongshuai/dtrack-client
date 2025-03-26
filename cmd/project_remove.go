/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"dtrack-client/types"
	"fmt"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a project by its UUID",
	Long: `Remove a project from the system using its unique identifier (UUID). This command permanently deletes the specified project and all associated data.

- Use the --id flag to specify the UUID of the project you want to remove. This flag is required.
  
Example usage:
- To remove a project with UUID "12345678-1234-1234-1234-1234567890ab":  
  dtrack project remove --id 12345678-1234-1234-1234-1234567890ab

Note: Once a project is removed, it cannot be recovered. Ensure that you have the correct UUID before executing this command.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("remove called")
		uuid := cmd.Flag("id").Value.String()
		defer Client.Close()
		res, err := Client.R().SetPathParam("id", uuid).Execute(types.ApiUrls[types.RemoveProject].Method, types.ApiUrls[types.RemoveProject].Url)
		if err != nil {
			fmt.Printf("删除项目失败:%s\n", err)
		}
		if !res.IsSuccess() {
			fmt.Printf("删除项目失败:%s\n", res.String())
			//return
		}
		//
		fmt.Printf("删除项目成功:%s\n", res.String())
	},
	Aliases: []string{"rm", "del", "delete"},
}

func init() {
	projectCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	removeCmd.Flags().String("id", "", "project uuid")
	removeCmd.MarkFlagRequired("id")
}
