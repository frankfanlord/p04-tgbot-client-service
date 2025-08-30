package telegram

// _bot default bot
var _bot Bot

// Bot 机器人
type Bot interface {
	// Initialize 初始化
	Initialize() error
	// Run 运行
	Run()
	// Shutdown 关闭
	Shutdown()
}

// ==============================================================================================================================================
// Bot 机器人 implement
// ==============================================================================================================================================

type bot struct {
	webhook     string        // the address that telegram send updates to
	secretToken string        // verify that the updates is legal
	channel     <-chan []byte // webhook transfer message to logic
	close       chan struct{}
	done        chan struct{}
}
