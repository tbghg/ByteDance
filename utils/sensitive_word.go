package utils

import (
	"ByteDance/pkg/common"
	"bufio"
	"github.com/syyongx/go-wordsfilter"
	"io"
	"os"
	"strconv"
)

//敏感词列表
var texts = make([]string, 250)

var wf = wordsfilter.New()
var root map[string]*wordsfilter.Node

func SensitiveWordInit() {
	file, err := os.Open(common.SensitiveWordsPath)
	if err != nil {
		Log.Error("文件打开错误" + err.Error())
		return
	}
	// 新建一个缓冲区，把内容先放在缓冲区
	reader := bufio.NewReader(file)
	i := 0
	// 循环读取文件中的内容，直到文件末尾位置
	for {
		// 遇到'\n'结束读取，但是'\n' 也读取进入
		buf, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF { // 文件已结束
				break
			}
			Log.Error("文件读取错误" + err.Error())
		}
		texts[i] = string(buf)
		i = i + 1
	}
	root = wf.Generate(texts)
}

// SensitiveWordCheck 敏感词检测
func SensitiveWordCheck(text string, userID int) bool {

	if len(texts) == 0 {
		// 文件读取失败，未使用敏感词检测
		return true
	}

	isContains := wf.Contains(text, root)
	if isContains {
		Log.Warn("UserID:" + strconv.Itoa(userID) + " | 发表: “" + text + "” | 被视为包含敏感词评论")
	}
	return isContains
}
