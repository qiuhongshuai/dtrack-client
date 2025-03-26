/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"dtrack-client/types"
	"encoding/json"
	"fmt"
	"resty.dev/v3"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload SBOM data to the specified project",
	Long: `Upload SBOM (Software Bill of Materials) data to a specified project in the system. This command supports both PUT and POST methods for uploading, with options to automatically create the project if it does not exist and to wait for the analysis to complete.

- Use the --id flag to specify the project UUID or the --name and --version flags to specify the project name and version.
- Provide the SBOM content either via the --file flag (path to the SBOM file) or the --content flag (raw SBOM content). The --file flag takes precedence if both are provided.
- Use the --async flag to control whether the command waits for the analysis to complete (default is true).
- Use the --autoCreate flag to enable automatic project creation if the specified project does not exist (default is true).

Example usage:
- To upload an SBOM file to a project with UUID "12345678-1234-1234-1234-1234567890ab":  
  dtrack sbom upload --id 12345678-1234-1234-1234-1234567890ab --file path/to/sbom.json
- To upload an SBOM file and wait for the analysis to complete:  
  dtrack sbom upload --name MyProject --version 1.0 --file path/to/sbom.json --async false

Note: Ensure that the project details and SBOM content are correctly provided. If the upload fails due to invalid data or other issues, an appropriate error message will be displayed.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if strings.TrimSpace(UploadArg.ProjectId) == "" && strings.TrimSpace(UploadArg.ProjectName) == "" {
			fmt.Println("请指定项目uuid或者项目名")
			return
		}
		if err := UploadArg.SetContent(); err != nil {
			fmt.Printf("设置sbom内容失败：%s", err)
			return
		}
		token := ""
		var res *resty.Response
		defer Client.Close()
		r := Client.R().SetURL(types.ApiUrls[types.UploadBom].Url).SetMethod(strings.ToUpper(UploadArg.Method))
		body := map[string]string{
			//"project":        UploadArg.ProjectId,
			"projectVersion": UploadArg.ProjectVersion,
			"projectName":    UploadArg.ProjectName,
			"autoCreate":     strconv.FormatBool(UploadArg.AutoCreate),
			"bom":            UploadArg.Content,
		}
		if strings.TrimSpace(UploadArg.ProjectId) != "" {
			body["project"] = UploadArg.ProjectId
		}
		switch strings.ToUpper(UploadArg.Method) {
		case "PUT":
			res, err = r.SetBody(body).Send()
		case "POST":
			res, err = r.SetMultipartFormData(body).Send()
		default:
			fmt.Println("上传方式错误")
			return
		}
		if err != nil {
			fmt.Printf("上传失败：%s", err)
			return
		}
		if !res.IsSuccess() {
			fmt.Printf("上传失败：%s", res.String())
			return
		}
		data := make(map[string]interface{})
		if err := json.Unmarshal(res.Bytes(), &data); err != nil {
			fmt.Printf("解析上传结果失败：%s", err)
			return
		}
		token, _ = data["token"].(string)
		if token != "" {
			fmt.Printf("上传成功，token为：%s\n", token)
		}
		if !UploadArg.Async {
			fmt.Println("上传成功，正在等待bom分析完成，请稍等...")
			for {
				res, err = Client.R().SetURL(types.ApiUrls[types.CheckProcess].Url).SetMethod(types.ApiUrls[types.CheckProcess].Method).
					SetPathParam("id", token).Send()
				if err != nil {
					fmt.Printf("查询bom分析状态失败：%s", err)
					return
					//break
				}
				if !res.IsSuccess() {
					fmt.Printf("查询bom分析状态失败：%s", res.String())
					return
					//break
				}
				r := make(map[string]interface{})
				if err := json.Unmarshal(res.Bytes(), &r); err != nil {
					fmt.Printf("解析bom分析状态失败：%s", err)
					return
				}
				//fmt.Println(res.String())
				if r["processing"].(bool) == false {
					fmt.Println("bom分析完成,请前往浏览器查看")
					break
				} else {
					fmt.Printf("bom分析中，请稍等...\n")
					time.Sleep(time.Second * 3)
				}
			}
		}
	},
}
var UploadArg = types.UploadBomArg{
	Method:     "PUT",
	Async:      true,
	AutoCreate: true,
}

func init() {
	sbomCmd.AddCommand(uploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	uploadCmd.Flags().StringVar(&UploadArg.ProjectId, "id", "", "project uuid")
	uploadCmd.Flags().StringVar(&UploadArg.ProjectName, "name", "", "project name")
	uploadCmd.Flags().StringVar(&UploadArg.ProjectVersion, "version", "", "project version")
	uploadCmd.MarkFlagsRequiredTogether("name", "version")
	uploadCmd.Flags().StringVar(&UploadArg.File, "file", "", "sbom文件路径，和content必须有一个填写，file优先于content")
	uploadCmd.Flags().StringVar(&UploadArg.Content, "content", "", "sbom文件内容，和file必须有一个填写，file优先于content")
	uploadCmd.MarkFlagsOneRequired("file", "content")
	uploadCmd.MarkFlagsMutuallyExclusive("file", "content")
	uploadCmd.Flags().BoolVar(&UploadArg.Async, "async", true, "是否异步上传，默认为true,如果为false，会一直等待直到bom分析完成")
	uploadCmd.Flags().BoolVar(&UploadArg.AutoCreate, "autoCreate", true, "是否自动创建项目，默认为true")
	uploadCmd.Flags().StringVar(&UploadArg.Method, "method", "PUT", "上传方式，默认为PUT，可选为PUT,POST")
}
