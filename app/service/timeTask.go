package service

import (
	"BscDapp/app/dao"
	"BscDapp/app/model"
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/util/gconv"
	"time"
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
	_, _ = gcron.AddSingleton("*/30 * * * * *", func() {
		g.Log().Debug("Task BscTask Begin")
		s.FinishBscTask()
		g.Log().Debug("Task BscTask End")
	}, "bscTask")

	//定时检测活动状态
	_, _ = gcron.AddSingleton("*/30 * * * * *", func() {
		g.Log().Debug("Task CheckGame Begin")
		_ = ListenTask.DealGameStatus(context.Background())
		g.Log().Debug("Task CheckGame End")
	}, "checkGame")
}

//读取队列任务
func (s *timeTask) FinishBscTask() {
	ctx := context.Background()
	task, err := dao.FaBscTask.Ctx(ctx).Where("status=0").Order("id asc").One()
	if err != nil {
		g.Log().Debug("Service TimeTask FinishBscTask Task Find Err:", err)
	}
	for task.Status == 0 {
		var (
			res    string
			err    error
			update g.Map
		)
		switch task.Type {
		case model.SendAddBalance:
			var param model.TaskAddUserBalance
			err = gconv.Struct(task.Task, &param)
			if err != nil {
				g.Log().Debug("Service TimeTask FinishBscTask SendAddBalance Struct Err:", err)
			}
			res, err = NewGame.AddUserBalance(&param)
		}
		if err != nil {
			update = g.Map{
				"remark":  err.Error(),
				"updated": time.Now().Unix(),
				"status":  2,
			}
		} else {
			update = g.Map{
				"remark":  res,
				"updated": time.Now().Unix(),
				"status":  1,
			}
		}
		_, err = dao.FaBscTask.Ctx(ctx).Where("id=?", task.Id).Update(update)
		if err != nil {
			g.Log().Debug("Service TimeTask FinishBscTask Task Update Err:", err)
			return
		}
		time.Sleep(time.Second)
		//查询新任务
		task, err = dao.FaBscTask.Ctx(ctx).Where("status=0").Order("id asc").One()
		if err != nil {
			g.Log().Debug("Service TimeTask FinishBscTask Task Find Err:", err)
		}
	}
}
