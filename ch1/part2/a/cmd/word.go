package cmd

import (
	"go-programming-tour-book/ch1/part2/a/internal/word"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamelCase
	ModeUnderscoreToLowerCamelCase
	ModeUnderscoreToUnderCase
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转化，模式如下：",
	"1:全部单词转为大写",
	"2:全部单词转为小写",
	"3:下划线单词转化为大写驼峰单词",
	"4:下划线单词转化为小写驼峰单词",
	"5:驼峰单词转下划线单词",
},
	"\n")

var str string
var mode int8

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转化",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUnderCase:
			content = word.CamelCaseToUnderscore(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		default:
			log.Fatalf("暂不支持该格式")
		}
		log.Printf("输出结果，%s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转化模式")
}
