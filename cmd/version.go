/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"dtrack-client/types"
	"dtrack-client/utils"
	"fmt"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the client and server version information",
	Long: `Retrieve and display the version information for both the client and the server. This command helps ensure compatibility between the CLI tool and the server by comparing their respective versions.

- The client version is retrieved directly from the CLI tool.
- The server version is fetched via an API request to the server. If the server is unreachable or an error occurs, the command will display an appropriate error message.

Example usage:
- To check the versions:  
  dtrack version

Output example:
	Client Version: 0.0.1
	Server Version: 0.0.1`,
	Run: func(cmd *cobra.Command, args []string) {
		clientVersion := rootCmd.Version
		var serverVersion string
		defer Client.Close()
		res, err := Client.R().Execute(types.ApiUrls[types.GetVersion].Method, types.ApiUrls[types.GetVersion].Url)
		if err != nil {
			serverVersion = err.Error()
		} else {
			r, err := utils.ParserResponse(res, "version")
			if err != nil {
				serverVersion = err.Error()
			} else {
				serverVersion = r["version"].(string)
			}
		}
		fmt.Printf("Client Version: %s\nServer Version: %s\n", clientVersion, serverVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
