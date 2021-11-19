package service

import (
	"BscDapp/app/chianService"
	"BscDapp/app/dao"
	"BscDapp/app/model"
	"context"
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"time"
)

var ListenTask = listenTask{}

type listenTask struct{}

//处理合约注册
func (s *listenTask) DealRegister(c context.Context, param *chianService.BscGameRegisterLog) error {
	return dao.FaBscUser.Transaction(c, func(ctx context.Context, tx *gdb.TX) error {
		refInfo, err := dao.FaBscUser.Ctx(ctx).Where("id=?", param.RefId).One()
		if err != nil || refInfo == nil {
			g.Log().Debug("Service DealRegister refInfo Find Err:", err)
			return gerror.New("推荐人不存在:" + err.Error())
		}
		var refStr string
		if refInfo.RefStr == "" {
			refStr = gconv.String(refInfo.Id)
		} else {
			refStr = gconv.String(refInfo.Id) + "," + refInfo.RefStr
		}
		insertInfo := model.FaBscUser{
			Id:      int(param.Id),
			Address: param.Addr.String(),
			RefStr:  refStr,
			RefId:   refInfo.Id,
			Created: int(time.Now().Unix()),
			Updated: int(time.Now().Unix()),
		}
		_, err = dao.FaBscUser.Ctx(ctx).OmitEmpty().Save(insertInfo)
		if err != nil {
			g.Log().Debug("Service DealRegister User Save Err:", err)
			return gerror.New("会员插入失败:" + err.Error())
		}
		//更新推荐人数和团队人数
		_, err = dao.FaBscUser.Ctx(ctx).Where("id=?", refInfo.Id).Increment("ref_num", 1)
		if err != nil {
			g.Log().Debug("Service DealRegister ref_num Increment Err:", err)
		}
		whereStr := fmt.Sprintf("id in (%s)", refStr)
		update := g.Map{
			"team_num": &gdb.Counter{
				Field: "team_num",
				Value: 1,
			},
			"updated": time.Now().Unix(),
		}
		_, err = dao.FaBscUser.Ctx(ctx).Where(whereStr).Update(update)
		if err != nil {
			g.Log().Debug("Service DealRegister team_num Increment Err:", err)
		}
		return nil
	})

}

//处理活动开关
func (s *listenTask) DealGameStatus(c context.Context) error {
	return dao.FaBscUserGame.Transaction(c, func(ctx context.Context, tx *gdb.TX) error {
		//查询当前进行中的活动
		gameInfo, err := dao.FaBscGameInfo.Ctx(ctx).Order("round desc").One()
		if err != nil {
			g.Log().Debug("Service Task DealGameStatus GameInfo Find Err:", err)
			return err
		}
		if gameInfo.Status == 1 {
			//检查是否到期
			if gameInfo.EndTime > int(time.Now().Unix()) {
				//没到期,退出检测
				return nil
			}
			//到期,关闭活动,发放奖励,添加交易任务,开启新活动
			_, err = dao.FaBscGameInfo.Ctx(ctx).Where("id=?", gameInfo.Id).Update(g.Map{"status": 2, "updated": time.Now().Unix()})
			if err != nil {
				g.Log().Debug("Service Task DealGameStatus GameInfo Close Err:", err)
				return err
			}

			//发放奖励给最后20人
			users, err := dao.FaBscUserGame.Ctx(ctx).Where("game_id=?", gameInfo.Id).Limit(20).All()
			if err != nil {
				g.Log().Debug("Service Task DealGameStatus UserGame All Err:", err)
				return err
			}
			//投资金额权重平分
			totalNum, err := dao.FaBscUserGame.Ctx(ctx).Where("game_id=?", gameInfo.Id).Limit(20).Sum("invest_num")
			if err != nil {
				g.Log().Debug("Service Task DealGameStatus totalNum Find Err:", err)
				return err
			}
			if len(users) > 0 && totalNum > 0 {
				for _, v := range users {
					//奖励会员并插入记录
					reward := gameInfo.JackPool * v.InvestNum / totalNum
					update := g.Map{
						"credit": &gdb.Counter{
							Field: "credit",
							Value: reward,
						},
						"total_credit": &gdb.Counter{
							Field: "total_credit",
							Value: reward,
						},
						"update": time.Now().Unix(),
					}
					_, err = dao.FaBscUser.Ctx(ctx).Where("id=?", v.Uid).Update(update)
					if err != nil {
						g.Log().Debug("Service Task DealGameStatus User Update Err:", err)
						return err
					}
					creditInfo := model.FaBscCredit{
						Uid:     v.Uid,
						Type:    model.CreditPool,
						Num:     reward,
						Created: int(time.Now().Unix()),
					}
					_, err = dao.FaBscCredit.Ctx(ctx).OmitEmpty().Save(creditInfo)
					if err != nil {
						g.Log().Debug("Service Task DealGameStatus Credit Save Err:", err)
						return err
					}
					//添加交易任务
					taskInfo := model.FaBscTask{
						Type:    model.SendAddBalance,
						Task:    gconv.String(model.AddUserBalance{Uid: v.Uid, Value: reward}),
						Updated: int(time.Now().Unix()),
					}
					_, err = dao.FaBscTask.Ctx(ctx).OmitEmpty().Save(taskInfo)
					if err != nil {
						g.Log().Debug("Service Task DealGameStatus Task Save Err:", err)
					}
				}
			}
		}
		//新开一场活动
		insertInfo := model.FaBscGameInfo{
			Round:    gameInfo.Round + 1,
			SeedPool: gameInfo.SeedPool * 50 / 100,
			JackPool: gameInfo.SeedPool * 50 / 100,
			EndTime:  int(time.Now().AddDate(0, 0, 1).Unix()), //24小时
			Status:   1,
			Updated:  int(time.Now().Unix()),
			Created:  int(time.Now().Unix()),
		}
		_, err = dao.FaBscGameInfo.Ctx(ctx).OmitEmpty().Save(insertInfo)
		if err != nil {
			g.Log().Debug("Service Task DealGameStatus GameInfo Save Err:", err)
			return err
		}

		return nil
	})
}

//处理会员购买门票
func (s *listenTask) DelBuyTicket(c context.Context, param *chianService.BscGameBuyTicketLog) error {
	return dao.FaBscUserTicket.Transaction(c, func(ctx context.Context, tx *gdb.TX) error {
		userInfo, err := dao.FaBscUser.Ctx(ctx).Where("address=?", param.Addr.String()).One()
		if err != nil || userInfo == nil {
			g.Log().Debug("Service DoListenTask DelBuyTicket UserInfo Find Err:", err)
			return gerror.New("获取会员信息失败：" + err.Error())
		}
		num := float64(param.Value.Uint64()) / float64(gconv.Float64(model.TokenDecimals))
		percent := float64(param.Percent) / float64(gconv.Float64(model.PercentDecimals))
		//插入购买记录
		buyInfo := model.FaBscUserTicket{
			Uid:       userInfo.Id,
			PayNum:    num,
			TicketNum: num * percent,
			Percent:   percent,
			Created:   int(time.Now().Unix()),
		}
		_, err = dao.FaBscUserTicket.OmitEmpty().Save(buyInfo)
		if err != nil || userInfo == nil {
			g.Log().Debug("Service DoListenTask DelBuyTicket UserTicket Save Err:", err)
			return gerror.New("会员购买门票信息保存失败：" + err.Error())
		}

		//注入奖池
		//jackNum:=num*25/100

		//注入种子池
		//seedNum:=num*25/100

		//奖励推荐人

		return nil
	})

}
