package cmd

import (
	"github.com/go-programming-tour-book/tour/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamelcase
	ModeUnderscoreToLowerCamelcase
	ModeCamelcaseToUnderscore
)

var desc = strings.Join([]string{
	"该命令支持各种单词格式转换， 模式如下：",
	"1: 全部转为大写",
	"2: 全部转为小写",
	"3：下划线单词转为大写驼峰单词",
	"4：下划线单词转为小写驼峰单词",
	"5：驼峰单词转为下划线单词",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "word format change",
	Long:  desc,
	Run:   func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelcase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelcase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelcaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换格式，请执行help word查看文档")
		}
		log.Printf("输出结果：%s", content)
	},
}
var str string
var mode int8

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换模式")
}
