package service

import (
	"BscDapp/app/dao"
	"BscDapp/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/util/gconv"
)

var TimeTask = timeTask{}

type timeTask struct{}

//循环监听
func (s *timeTask) ListenTask() {
	//循环监听BSC日志
	_, _ = gcron.AddSingleton("1 */1 * * * *", func() {
		//监听之前先补全为监听到的日志
		baseInfo, _ := dao.FaBscBaseInfo.Where("theKey=?", model.BaseReadKey).One()
		NewGame.ReadBlockLog(gconv.Int64(baseInfo.TheValue))

		//开始监听
		NewGame.ListenNewGame()

		//监听报错,初始化服务
		NewGame.WsClient.Close()
		NewGame.Client.Close()
		NewGame.Init()
	}, "bscListen")

	//定时消费任务
	//_, _ = gcron.AddSingleton("*/30 * * * * *", func() {
	//	g.Log().Debug("Task BscTask Begin")
	//	s.FinishBscTask()
	//	g.Log().Debug("Task BscTask End")
	//}, "bscTask")

	//定时检测活动状态
	//_, _ = gcron.AddSingleton("*/5 * * * * *", func() {
	//	g.Log().Debug("Task CheckGame Begin")
	//	_ = ListenTask.DealGameStatus(context.Background())
	//	g.Log().Debug("Task CheckGame End")
	//}, "checkGame")
}

//读取队列任务
func (s *timeTask) FinishBscTask() {
	g.Log().Debug("消费一条任务")
}
