package service

import (
	"BscDapp/app/dao"
	"BscDapp/app/model"
	"context"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/util/gconv"
	"strconv"
	"time"
)

var TimeTask = timeTask{}

type timeTask struct{}

//定时任务
func (s *timeTask) ListenTask() {
	//循环监听BSC日志
	_, _ = gcron.AddSingleton("1 */1 * * * *", func() {
		//监听之前先补全为监听到的日志
		baseInfo, _ := dao.FaBscBaseInfo.Where("theKey=?", model.BaseReadKey).One()
		num, err := strconv.Atoi(baseInfo.TheValue)
		if err != nil {
			num = 1
			g.Log().Debug("Service ListenTask strconv Err:", err)
		}
		NewGame.ReadBlockLog(int64(num))
		//开始监听
		NewGame.ListenNewGame()

		//监听报错,初始化服务
		NewGame.WsClient.Close()
		NewGame.Client.Close()
		NewGame.Init()
	}, "bscListen")
	//定时检测活动状态
	_, _ = gcron.AddSingleton("1 */10 * * * *", func() {
		g.Log().Debug("Task CheckGame Begin")
		_ = ListenTask.DealGameStatus(context.Background())
		g.Log().Debug("Task CheckGame End")
	}, "checkGame")

	//定时消费任务
	_, _ = gcron.AddSingleton("1 */11 * * * *", func() {
		g.Log().Debug("Task BscTask Begin")
		s.FinishBscTask()
		g.Log().Debug("Task BscTask End")
	}, "bscTask")
}

//读取队列任务
func (s *timeTask) FinishBscTask() {
	ctx := context.Background()
	task, err := dao.FaBscTask.Ctx(ctx).Where("status=0 and task_time<?", time.Now().Unix()).Order("task_time asc").One()
	if err != nil {
		g.Log().Debug("Service TimeTask FinishBscTask Task Find Err:", err)
		return
	}
	for task != nil {
		var (
			res    string
			err    error
			update = g.Map{"updated": time.Now().Unix()}
		)
		switch task.Type {
		case model.SendPay:
			var param model.TaskAddUserBalance
			err = gconv.Struct(task.Task, &param)
			if err != nil {
				g.Log().Debug("Service TimeTask FinishBscTask SendPay Struct Err:", err)
			}
			res, err = NewGame.Pay(&param)
		case model.SendSetRound:
			res, err = NewGame.SetRound()
		case model.SendUserOut:
			var param model.TaskUserOut
			err = gconv.Struct(task.Task, &param)
			if err != nil {
				g.Log().Debug("Service TimeTask FinishBscTask SendAddBalance Struct Err:", err)
			}
			res, err = NewGame.UserOut(&param)
		}

		if err != nil {
			update["remark"] = err.Error()
			update["task_num"] = task.TaskNum + 1
			update["task_time"] = time.Now().Unix() + int64((task.TaskNum+1)*model.TaskLateTime)
		} else {
			update["remark"] = res
			update["status"] = 1
		}
		_, err = dao.FaBscTask.Ctx(ctx).Where("id=?", task.Id).Update(update)
		if err != nil {
			g.Log().Debug("Service TimeTask FinishBscTask Task Update Err:", err)
			return
		}
		time.Sleep(time.Second)
		//查询新任务
		task, err = dao.FaBscTask.Ctx(ctx).Where("status=0 and task_time<?", time.Now().Unix()).Order("task_time asc").One()
		if err != nil {
			g.Log().Debug("Service TimeTask FinishBscTask Task Find Err:", err)
		}
	}
}
