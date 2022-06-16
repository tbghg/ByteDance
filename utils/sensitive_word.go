package utils

import "github.com/syyongx/go-wordsfilter"

//敏感词列表
var texts = []string{
	"涉黄",
	"涉政。。。。。",
	"邪教。。。。。",
	"欺骗。。。。。",
}

//敏感词检测
func SensitiveWordCheck(text string) bool {

	wf := wordsfilter.New()

	// Generate
	root := wf.Generate(texts)
	// Generate with file
	// root := wf.GenerateWithFile(path)

	// Contains
	c1 := wf.Contains(text, root)
	//涉及敏感词
	if c1 {
		return true
	}
	return false
}
