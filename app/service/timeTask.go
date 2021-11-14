package service

import "github.com/gogf/gf/os/gcron"

var TimeTask = timeTask{}

type timeTask struct{}

//定时任务
func (s *timeTask) CompleteTask() {
	//循环监听BSC日志
	_, _ = gcron.AddSingleton("* * * * * *", func() {
		NewGame.ListenNewGame()
		NewGame.WsClient.Close()
		NewGame.Client.Close()
		NewGame.Init()
	})

}
