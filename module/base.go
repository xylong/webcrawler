package module

type Counts struct {
	// CalledCount 调用计数
	CalledCount uint64
	// AcceptedCount 接受计数
	AcceptedCount uint64
	// CompletedCount 成功完成计数
	CompletedCount uint64
	// HandlingNumber 实时处理数
	HandlingNumber uint64
}

// Module 代表组件的基础接口类型
type Module interface {
	// ID 获取当前组件的ID
	ID() MID
	// Addr 获取当前组件的网络地址
	Addr() string
	// Score 获取当前组件的评分
	Score() uint64
	// SetScore 设置当前组件的评分
	SetScore(score uint64)
	// ScoreCalculator 获取评分计算器
	ScoreCalculator() CalculateScore
	// CallCount 获取当前组件被调用的计数器
	CalledCount() uint64
	// AcceptedCount 获取被当前组件接受的调用的计数
	AcceptedCount() uint64
	// CompletedCount 获取当前组件已成功完成的调用的计数。
	CompletedCount() uint64
	// HandlingNumber 获取当前组件正在处理的调用的数量。
	HandlingNumber() uint64
	//Counts 一次性获取所有计数。
	Counts()
	// Summary 获取组件摘要
	Summary() Summary
}

// Summary 组件摘要
type Summary struct {
	ID        MID         `json:"id"`
	Called    uint64      `json:"called"`
	Accepted  uint64      `json:"accepted"`
	Completed uint64      `json:"completed"`
	Handling  uint64      `json:"handling"`
	Extra     interface{} `json:"extra,omitempty"`
}
