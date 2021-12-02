package service

import (
	"BscDapp/app/dao"
	"BscDapp/app/model"
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"time"
)

var User = user{}

type user struct{}

//获取登陆会员信息
func (s *user) GetUser(r *ghttp.Request) *model.FaBscUser {
	user := new(model.FaBscUser)
	data := r.GetCtxVar("user")
	if err := gconv.Struct(data, user); err != nil {
		g.Log().Debug("Service User GetUser Err:", err)
		return nil
	}
	return user
}

//修改会员余额
func (s *user) ChangeCredit(ctx context.Context, uid int, amount float64, doType string) error {
	var (
		update   = g.Map{}
		err      error
		filename = model.UserCreditOne
	)

	switch doType {
	case model.CreditPool: //奖池奖励
		update = g.Map{
			"pool_credit": &gdb.Counter{
				Field: "pool_credit",
				Value: amount,
			},
			"total_credit": &gdb.Counter{
				Field: "total_credit",
				Value: amount,
			},
			"updated": time.Now().Unix(),
		}
		filename = model.UserCreditThree
	case model.CreditRefReward: //推荐奖励
		update = g.Map{
			"credit": &gdb.Counter{
				Field: "credit",
				Value: amount,
			},
			"total_credit": &gdb.Counter{
				Field: "total_credit",
				Value: amount,
			},
			"updated": time.Now().Unix(),
		}
	case model.CreditReward: //投资回报
		update = g.Map{
			"credit": &gdb.Counter{
				Field: "credit",
				Value: amount,
			},
			"total_credit": &gdb.Counter{
				Field: "total_credit",
				Value: amount,
			},
			"updated": time.Now().Unix(),
		}
	case model.CreditPoolWithdraw: //奖池提现
		update = g.Map{
			"pool_credit": &gdb.Counter{
				Field: "pool_credit",
				Value: amount,
			},
			"updated": time.Now().Unix(),
		}
		//添加任务发放到合约
		taskInfo := model.FaBscTask{
			Type:     model.SendPay,
			Task:     gconv.String(model.TaskAddUserBalance{UserId: uint64(uid), Value: -amount}),
			Created:  int(time.Now().Unix()),
			TaskTime: int(time.Now().Unix()),
		}
		_, err = dao.FaBscTask.Ctx(ctx).OmitEmpty().Save(taskInfo)
		if err != nil {
			g.Log().Debug("Service User ChangeCredit Task Save Err:", err)
			return err
		}
		filename = model.UserCreditThree
	case model.CreditWithdraw: //投资回报提现
		update = g.Map{
			"credit": &gdb.Counter{
				Field: "credit",
				Value: amount,
			},
			"updated": time.Now().Unix(),
		}
		//添加任务发放到合约
		taskInfo := model.FaBscTask{
			Type:     model.SendPay,
			Task:     gconv.String(model.TaskAddUserBalance{UserId: uint64(uid), Value: -amount}),
			Created:  int(time.Now().Unix()),
			TaskTime: int(time.Now().Unix()),
		}
		_, err = dao.FaBscTask.Ctx(ctx).OmitEmpty().Save(taskInfo)
		if err != nil {
			g.Log().Debug("Service User ChangeCredit Task Save Err:", err)
			return err
		}
	case model.CreditBuyTicket: //购买门票
		update = g.Map{
			"ticket_num": &gdb.Counter{
				Field: "ticket_num",
				Value: amount,
			},
			"updated": time.Now().Unix(),
		}
		filename = model.UserCreditTwo
	default:
		return gerror.New("Service User ChangeCredit Err:会员余额事件不存在")
	}
	_, err = dao.FaBscUser.Ctx(ctx).Where("id=?", uid).Update(update)
	if err != nil {
		g.Log().Debug("Service User ChangeCredit User Update Err:", err)
		return err
	}
	creditInfo := model.FaBscCredit{
		Uid:      uid,
		FileName: filename,
		Type:     doType,
		Num:      amount,
		Created:  int(time.Now().Unix()),
	}
	_, err = dao.FaBscCredit.Ctx(ctx).OmitEmpty().Save(creditInfo)
	if err != nil {
		g.Log().Debug("Service User ChangeCredit Credit Save Err:", err)
	}
	return err
}
