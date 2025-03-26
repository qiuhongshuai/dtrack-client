/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "设置配置文件的相关参数",
	Long:  `设置配置文件的相关参数`,
	Run: func(cmd *cobra.Command, args []string) {
		server := cmd.Flag("server").Value.String()
		apikey := cmd.Flag("apikey").Value.String()
		//team := cmd.Flag("team").Value.String()
		timeout := cmd.Flag("timeout").Value.String()
		//if server == "" && apikey == "" && team == "" && timeout == "0" {
		//	fmt.Println("no params")
		//	return
		//}
		if server != "" {
			viper.Set("server", server)

		}
		if apikey != "" {
			viper.Set("apikey", apikey)
		}
		//if team != "" {
		//	viper.Set("team", team)
		//}
		if timeout != "0" {
			tw, err := strconv.Atoi(timeout)
			if err == nil {
				viper.Set("timeout", tw)
			}
		}
		if err := viper.WriteConfig(); err != nil {
			fmt.Printf("set config error:%s", err.Error())
		}
	},
}

func init() {
	configCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	setCmd.Flags().String("server", "", "设置dtrack地址")
	setCmd.Flags().String("apikey", "", "设置dtrack认证的apiKey")
	//setCmd.Flags().String("team", "", "set team")
	setCmd.Flags().Int("timeout", 5, "设置请求服务端的超时时间")
}
