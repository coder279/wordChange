package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strings"
	"wordChange/internal/word"
)

const (
	ModeUpper                      = iota + 1 //全部单词转为大写
	ModeLower                                 //全部单词转为小写
	ModeUnderscoreToUpperCamelcase            //下划线单词转为大写驼峰单词
	ModeUnderscoreToLowerCamelcase            //下划线单词转为小写驼峰单词
	ModeCamelcaseToUnderscore                 //驼峰单词转为下划线单词
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换,模式如下: ",
	"1: 全部单词转为大写",
	"2: 全部单词转为小写",
	"3: 下划线单词转为大写驼峰单词",
	"4: 下划线单词转为小写驼峰单词",
	"5: 驼峰单词转为下划线单词",
},"\n")

var (
	str string
	mode int8
)
var wordCmd = &cobra.Command{
	Use:  "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeCamelcaseToUnderscore:
			content = word.CamelCaseToUnderScore(str)
		case ModeUnderscoreToUpperCamelcase:
			content = word.UnderScoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelcase:
			content = word.UnderScoreToLowerCamelCase(str)
		default:
			log.Fatalf("暂不支持该种转换方式,请执行help word查看帮助")
		}
		log.Printf("输出结果: %s",content)
	},
}
func init(){
	wordCmd.Flags().StringVarP(&str,"str","s","","请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode,"mode","m",0,"请输入单词转换模式")
}