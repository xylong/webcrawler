package module

// Type 代表组件的类型。
type Type string

// 当前认可的组件类型的常量
const (
	// TYPE_DOWNLOADER 下载器
	TYPE_DOWNLOADER Type = "downloader"
	// TYPE_ANALYZER 分析器
	TYPE_ANALYZER Type = "analyzer"
	// TYPE_PIPELINE 条目处理管道
	TYPE_PIPELINE Type = "pipeline"
)

// legalTypeLetterMap 代表合法的组件类型-字母的映射。
var legalTypeLetterMap = map[Type]string{
	TYPE_DOWNLOADER: "D",
	TYPE_ANALYZER:   "A",
	TYPE_PIPELINE:   "P",
}
