package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
)

var TimeTask = timeTask{}

type timeTask struct{}

//循环监听
func (s *timeTask) ListenTask() {
	//循环监听BSC日志
	_, _ = gcron.AddSingleton("1 */1 * * * *", func() {
		NewGame.ListenNewGame()
		NewGame.WsClient.Close()
		NewGame.Client.Close()
		NewGame.Init()
	}, "bscListen")

	//定时消费任务
	_, _ = gcron.AddSingleton("*/30 * * * * *", func() {
		s.FinishBscTask()
	}, "bscTask")

	//定时检测活动状态
	_, _ = gcron.AddSingleton("*/5 * * * * *", func() {
	}, "checkGame")
}

//监听活动是否结束
func CheckGameStatus() {

}

//读取队列任务
func (s *timeTask) FinishBscTask() {
	g.Log().Debug("消费一条任务")
}
