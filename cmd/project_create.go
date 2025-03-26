/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"dtrack-client/types"
	"dtrack-client/utils"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project with specified details",
	Long: `Create a new project by providing necessary details such as name, version, classifier, and optional parent ID. The command automatically assigns the current team to the project's access list.

- Use the --parent-id flag to specify the UUID of the parent project (optional).
- Use the --classifier flag to define the type of the project (default is APPLICATION). Available classifiers include OPERATING_SYSTEM, APPLICATION, LIBRARY, FRAMEWORK, CONTAINER, DEVICE, FILE, FIRMWARE, and NONE.
- If no version is provided, the default value "latest" will be used, marking the project as the latest version.

The command retrieves the current team information to associate it with the project. If any required field is missing or an error occurs during execution, the command will provide an appropriate error message.

Example usage:
- To create a project with default settings: dtrack project create --name "MyProject" --version "1.0"
- To create a project with a specific classifier: dtrack project create --name "MyLibrary" --version "1.0" --classifier LIBRARY

Note: Ensure that the classifier value is valid and matches one of the predefined options.`,
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Parent().Flag("name").Value.String()
		if name == "" {
			fmt.Printf("请指定项目名称")
			return
		}
		version := cmd.Parent().Flag("version").Value.String()
		if version == "" {
			version = "latest"
		}
		parentId := cmd.Flag("parent-id").Value.String()
		classifier := strings.ToUpper(cmd.Flag("classifier").Value.String())
		if !utils.Contains(classifiers, classifier) {
			fmt.Printf("分类器错误,可选值：【%s】", strings.Join(classifiers, ","))
			return
		}
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
		var project types.Project
		project.Name = name
		project.Version = version
		project.Classifier = classifier

		if project.Version == "latest" {
			project.IsLatest = true
		}
		if parentId != "" {
			project.Parent.Uuid = parentId
		}
		project.Active = true
		project.AccessTeams = append(project.AccessTeams, &team)
		//fmt.Printf("正在创建项目%+v\n", team)
		res, err = Client.R().SetContentType("application/json").SetBody(project).Execute(types.ApiUrls[types.CreateProject].Method, types.ApiUrls[types.CreateProject].Url)
		if err != nil {
			fmt.Printf("创建项目报错：%s", err.Error())
			return
		}
		if !res.IsSuccess() {
			//fmt.Println(res.Request.CurlCmd())
			fmt.Printf("创建项目失败：%s", res.String())
			return
		}
		//fmt.Println(res.Request.CurlCmd())
		fmt.Printf("创建项目%s成功", name)
	},
}

var classifiers = []string{"OPERATING_SYSTEM", "APPLICATION", "LIBRARY", "FRAMEWORK", "CONTAINER", "DEVICE", "FILE", "FIRMWARE", "NONE"}

func init() {
	projectCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().String("parent-id", "", "父项目ID，可以为空")
	createCmd.Flags().String("classifier", "APPLICATION", fmt.Sprintf("分类器，可选值：【%s】", strings.Join(classifiers, ",")))
	createCmd.Flag("classifier").NoOptDefVal = "APPLICATION"
	createCmd.Flag("classifier").Annotations = map[string][]string{
		cobra.BashCompCustom: {"__cobra_comp_classifier"},
	}
	createCmd.RegisterFlagCompletionFunc("classifier", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return classifiers, cobra.ShellCompDirectiveDefault
	})
}
