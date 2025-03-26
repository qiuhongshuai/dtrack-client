/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "project管理",
	Long:  `project管理`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("project called")
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	projectCmd.PersistentFlags().StringP("name", "n", "", "project name")
	projectCmd.PersistentFlags().StringP("version", "V", "", "project 版本")
}
