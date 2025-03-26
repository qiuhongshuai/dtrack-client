/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// sbomCmd represents the sbom command
var sbomCmd = &cobra.Command{
	Use:   "sbom",
	Short: "Manage Software Bill of Materials (SBOM) operations",
	Long:  "Manage Software Bill of Materials (SBOM) operations",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(sbomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sbomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sbomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
