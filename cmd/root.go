/*
Copyright © 2025 hessqiu<qiuhs1@dazd.cn>

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
package cmd

import (
	"crypto/tls"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"resty.dev/v3"
	"time"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dtrack-client",
	Short: "Dependency Track客户端工具",
	Long: `Dependency Track 客户端工具是一个用于与 Dependency Track 服务器交互的实用程序。
它可以帮助用户将软件物料清单（SBOM）上传到 Dependency Track 服务器。通过该工具，
用户可以轻松地在 CI/CD 流程中集成 Dependency Track 的功能，实现对项目依赖组件的安全性和合规性检查`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("root called")
	//	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
	//		fmt.Printf("%s=%s\n", flag.Name, flag.Value.String())
	//	})
	//	fmt.Println("Server:", viper.Get("server"))
	//},
	Version: "v0.0.1",
	//CompletionOptions: cobra.CompletionOptions{
	//	DisableDefaultCmd: true,
	//},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.

var Client *resty.Client

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dtrack-client.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringP("config", "C", ".dtrack.yaml", "配置文件位置")
	viper.SetConfigFile(rootCmd.Flag("config").Value.String())
	_, err := os.Stat(rootCmd.Flag("config").Value.String())
	if os.IsNotExist(err) {
		_, err := os.Create(rootCmd.Flag("config").Value.String())
		if err != nil {
			fmt.Println("创建配置文件失败")
			os.Exit(1)
		}
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("读取配置文件失败")
		os.Exit(1)
	}
	rootCmd.PersistentFlags().StringP("server", "S", "http://localhost", "dtrack地址")
	viper.BindPFlag("server", rootCmd.PersistentFlags().Lookup("server"))
	rootCmd.PersistentFlags().StringP("apikey", "A", "", "dtrack 认证使用的apikey")
	viper.BindPFlag("apikey", rootCmd.PersistentFlags().Lookup("apikey"))
	rootCmd.PersistentFlags().BoolP("skip-verify", "k", false, "是否跳过https证书认证，只有在使用https时生效")
	viper.BindPFlag("skip-verify", rootCmd.PersistentFlags().Lookup("skip-verify"))
	rootCmd.PersistentFlags().IntP("timeout", "t", 5, "请求超时时间")
	viper.BindPFlag("timeout", rootCmd.PersistentFlags().Lookup("timeout"))
	Client = resty.New()
	Client.SetBaseURL(viper.GetString("server"))
	if viper.GetBool("skip-verify") {
		Client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	Client.SetHeaderAuthorizationKey("X-API-Key")
	Client.SetAuthToken(viper.GetString("apikey"))
	Client.SetAuthScheme("")
	Client.SetTimeout(time.Duration(viper.GetInt("timeout")) * time.Second)
	//fmt.Println(Client.BaseURL(), "abbc", viper.Get("server"))
}
