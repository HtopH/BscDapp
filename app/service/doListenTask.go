package service

import (
	"BscDapp/app/chainService"
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
func (s *listenTask) DealRegister(c context.Context, param *chainService.BscGameRegisterLog) error {
	return dao.FaBscUser.Transaction(c, func(ctx context.Context, tx *gdb.TX) error {
		//校验参数
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
		gameInfo, err := dao.FaBscGameInfo.Ctx(ctx).Where("status=1").Order("round desc").One()
		if err != nil {
			g.Log().Debug("Service Task DealGameStatus GameInfo Find Err:", err)
			return err
		}
		if gameInfo == nil {
			gameInfo = &model.FaBscGameInfo{}
		}
		if gameInfo != nil && gameInfo.Status == 1 {
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
			userGames, err := dao.FaBscUserGame.Ctx(ctx).Where("game_round=?", gameInfo.Round).Limit(20).Order("created desc").All()
			if err != nil {
				g.Log().Debug("Service Task DealGameStatus UserGame All Err:", err)
				return err
			}
			//投资金额权重平分
			totalNum, err := dao.FaBscUserGame.Ctx(ctx).Where("game_round=?", gameInfo.Round).Limit(20).Order("created desc").Sum("invest_num")
			if err != nil {
				g.Log().Debug("Service Task DealGameStatus totalNum Find Err:", err)
				return err
			}
			if len(userGames) > 0 && totalNum > 0 {
				for _, v := range userGames {
					//奖励会员并插入记录
					reward := gameInfo.JackPool * v.InvestNum / totalNum
					err = User.ChangeCredit(ctx, v.Uid, reward, model.CreditPool)
					if err != nil {
						return err
					}
					//更新奖励金额
					_, err = dao.FaBscUserGame.Ctx(ctx).Where("id=?", v.Id).Update(g.Map{"award_num": reward})
					if err != nil {
						g.Log().Debug("Service Task DealGameStatus UserGame Update Err:", err)
					}
				}
			}
		}
		//新开一场活动
		insertInfo := model.FaBscGameInfo{
			Round:    gameInfo.Round + 1,
			SeedPool: gameInfo.SeedPool * 50 / 100,
			JackPool: gameInfo.SeedPool * 50 / 100,
			EndTime:  int(time.Now().Unix() + model.GameLevelOneTime), //24小时
			Status:   1,
			Updated:  int(time.Now().Unix()),
			Created:  int(time.Now().Unix()),
		}
		g.Log().Println(insertInfo)
		_, err = dao.FaBscGameInfo.Ctx(ctx).OmitEmpty().Save(insertInfo)
		if err != nil {
			g.Log().Debug("Service Task DealGameStatus GameInfo Save Err:", err)
			return err
		}
		//更新合约场次,添加交易任务
		taskInfo := model.FaBscTask{
			Type:     model.SendSetRound,
			Created:  int(time.Now().Unix()),
			TaskTime: int(time.Now().Unix()),
		}
		_, err = dao.FaBscTask.Ctx(ctx).OmitEmpty().Save(taskInfo)
		if err != nil {
			g.Log().Debug("Service Task DealGameStatus Task Save Err:", err)
		}
		return err
	})
}

//处理会员购买门票
func (s *listenTask) DealBuyTicket(c context.Context, param *chainService.BscGameBuyTicketLog) error {
	return dao.FaBscUserTicket.Transaction(c, func(ctx context.Context, tx *gdb.TX) error {
		userInfo, err := dao.FaBscUser.Ctx(ctx).Where("id=?", param.Id).One()
		if err != nil || userInfo == nil {
			g.Log().Debug("Service DoListenTask DelBuyTicket UserInfo Find Err:", err)
			return gerror.New("获取会员信息失败：" + err.Error())
		}
		num := BigIntToF(param.Value, model.TokenDecimals)
		percent := float64(param.Percent) / gconv.Float64(model.PercentDecimals)
		ticketNum := num * percent
		//插入购买记录
		buyInfo := model.FaBscUserTicket{
			Uid:       userInfo.Id,
			PayNum:    num,
			TicketNum: ticketNum,
			Percent:   percent,
			Created:   int(time.Now().Unix()),
		}
		_, err = dao.FaBscUserTicket.Ctx(ctx).OmitEmpty().Save(buyInfo)
		if err != nil || userInfo == nil {
			g.Log().Debug("Service DoListenTask DelBuyTicket UserTicket Save Err:", err)
			return gerror.New("会员购买门票信息保存失败：" + err.Error())
		}
		err = User.ChangeCredit(ctx, userInfo.Id, ticketNum, model.CreditBuyTicket)
		if err != nil {
			return err
		}
		//注入奖池
		jackNum := num * 25 / 100
		//注入种子池
		seedNum := num * 25 / 100
		gameInfo, err := dao.FaBscGameInfo.Ctx(ctx).Where("status=1").One()
		if err != nil || gameInfo == nil {
			g.Log().Debug("Service DoListenTask DelBuyTicket GameInfo One Err:", err)
			return gerror.New("活动信息查询失败：" + err.Error())
		}
		updateGame := g.Map{
			"seed_pool": &gdb.Counter{
				Field: "seed_pool",
				Value: seedNum,
			},
			"jack_pool": &gdb.Counter{
				Field: "jack_pool",
				Value: jackNum,
			},
			"updated": time.Now().Unix(),
		}
		//判断奖池是否触发更新结束时间
		endTime := GetGameEndTime(gameInfo, gameInfo.JackPool+jackNum)
		if endTime > 0 {
			updateGame["end_time"] = endTime
		}
		_, err = dao.FaBscGameInfo.Ctx(ctx).Where("id=?", gameInfo.Id).Update(updateGame)
		if err != nil || gameInfo == nil {
			g.Log().Debug("Service DoListenTask DelBuyTicket GameInfo Update Err:", err)
			return gerror.New("活动奖池更新失败：" + err.Error())
		}
		//奖励推荐人
		if userInfo.RefId > 0 {
			err = User.ChangeCredit(ctx, userInfo.RefId, num*20/100, model.CreditRefReward)
			if err != nil {
				return err
			}
		}

		//更新系统已兑换门票
		spendInfo, _ := dao.FaBscBaseInfo.Ctx(ctx).Where("theKey=?", model.BaseSpendKey).One()
		if spendInfo != nil {
			spend := ticketNum + gconv.Float64(spendInfo.TheValue)
			_, _ = dao.FaBscBaseInfo.Ctx(ctx).Where("theKey=?", model.BaseSpendKey).Update(g.Map{"theValue": spend, "updated": time.Now().Unix()})
		}
		return nil
	})

}

//处理会员投资
func (s *listenTask) DealUserJoinGame(c context.Context, param *chainService.BscGameJoinLog) error {
	return dao.FaBscUserGame.Transaction(c, func(ctx context.Context, tx *gdb.TX) error {
		userInfo, err := dao.FaBscUser.Ctx(ctx).Where("id=?", param.Id).One()
		if err != nil || userInfo == nil {
			g.Log().Debug("Service DoListenTask DealUserJoinGame UserInfo Find Err:", err)
			return gerror.New("获取会员信息失败：" + err.Error())
		}
		num := BigIntToF(param.Value, model.TokenDecimals)
		ticket := num * model.PercentJoinTicket / model.PercentBase
		gameInfo, err := dao.FaBscGameInfo.Ctx(ctx).Where("status=1 and round=?", param.Round).One()
		if err != nil || gameInfo == nil {
			g.Log().Debug("Service DoListenTask DealUserJoinGame gameInfo Find Err:", err)
			return gerror.New("查询活动信息失败:" + err.Error())
		}
		now := time.Now()
		insertUserGame := model.FaBscUserGame{
			Uid:       userInfo.Id,
			GameRound: gameInfo.Round,
			InvestNum: num,
			TicketNum: ticket,
			ReturnNum: num * model.PercentJoinReturn / model.PercentBase,
			Status:    1,
			Created:   int(now.Unix()),
		}

		_, err = dao.FaBscUserGame.Ctx(ctx).OmitEmpty().Save(insertUserGame)
		if err != nil || gameInfo == nil {
			g.Log().Debug("Service DoListenTask DealUserJoinGame UserGame Save Err:", err)
			return gerror.New("插入活动参与记录失败:" + err.Error())
		}
		//更新会员投资金额
		_, err = dao.FaBscUser.Ctx(ctx).Where("id=?", userInfo.Id).Increment(dao.FaBscUser.Columns.InvestNum, num)
		if err != nil {
			g.Log().Debug("Service DoListenTask DealUserJoinGame User Increment Err:", err)
			return gerror.New("更新会员投资金额失败:" + err.Error())
		}

		//会员参与活动,更新活动信息
		nowJack := gameInfo.JackPool + num*model.PercentJoinToJack/model.PercentBase
		nowSeed := gameInfo.JackPool + num*model.PercentJoinToSeed/model.PercentBase
		endTime := GetGameEndTime(gameInfo, nowJack)
		if endTime == 0 {
			//结束时间加10分钟
			endTime = gameInfo.EndTime + model.GameBaseTime
		}
		updateGame := g.Map{
			"jack_pool": nowJack,
			"seed_pool": nowSeed,
			"end_time":  endTime,
			"updated":   now.Unix(),
		}
		_, err = dao.FaBscGameInfo.Ctx(ctx).Where("id=?", gameInfo.Id).Update(updateGame)
		if err != nil || gameInfo == nil {
			g.Log().Debug("Service DoListenTask DealUserJoinGame GameInfo Update Err:", err)
			return gerror.New("更新活动信息失败:" + err.Error())
		}
		//70%反馈给前面投资者
		reward := num * model.PercentJoinOut / model.PercentBase
		userGame, err := dao.FaBscUserGame.Ctx(ctx).Where("game_round=? and return_num>0 and status=1", gameInfo.Round).Order("created asc").One()
		if err != nil {
			g.Log().Debug("Service DoListenTask DealUserJoinGame UserGame Find Err:", err)
			return gerror.New("查询会员活动信息失败:" + err.Error())
		}
		for userGame != nil && reward > 0 {
			var canGet float64
			var updateUserGame = g.Map{}
			//计算回报金额,判断是否出局
			if userGame.ReturnNum > reward {
				canGet = reward
				reward = 0
			} else {
				canGet = userGame.ReturnNum
				reward = reward - userGame.ReturnNum
				//满足出局条件
				updateUserGame["status"] = 2
				//添加任务通知合约会员出局
				taskInfo := model.FaBscTask{
					Type:     model.SendUserOut,
					Task:     gconv.String(model.TaskUserOut{UserId: uint64(userGame.Uid), Round: uint32(gameInfo.Round)}),
					Created:  int(time.Now().Unix()),
					TaskTime: int(time.Now().Unix()),
				}
				_, err = dao.FaBscTask.Ctx(ctx).OmitEmpty().Save(taskInfo)
				if err != nil {
					g.Log().Debug("Service DoListenTask DealUserJoinGame Task Save Err:", err)
					return err
				}
			}
			updateUserGame["return_num"] = userGame.ReturnNum - canGet
			_, err = dao.FaBscUserGame.Ctx(ctx).Where("id=?", userGame.Id).Update(updateUserGame)
			if err != nil {
				g.Log().Debug("Service DoListenTask DealUserJoinGame userGame Update Err:", err)
				return err
			}
			//发放到账户
			err = User.ChangeCredit(ctx, userGame.Uid, canGet, model.CreditReward)

			if reward > 0 {
				//查询下一个记录
				userGame, err = dao.FaBscUserGame.Ctx(ctx).Where("game_round=? and return_num>0 and status=1", gameInfo.Round).Order("created asc").One()
				if err != nil {
					g.Log().Debug("Service DoListenTask DealUserJoinGame UserGame Find Err:", err)
					return gerror.New("查询会员活动信息失败:" + err.Error())
				}
			}
		}
		return nil
	})

}

//处理会员提现
func (s *listenTask) DealUserWithdraw(c context.Context, param *chainService.BscGameUserGetLog) error {
	return dao.FaBscUser.Transaction(c, func(ctx context.Context, tx *gdb.TX) error {
		userInfo, err := dao.FaBscUser.Ctx(ctx).Where("id=?", param.Id).One()
		if err != nil || userInfo == nil {
			g.Log().Debug("Service DoListenTask DealUserWithdraw UserInfo Find Err:", err)
			return gerror.New("获取会员信息失败：" + err.Error())
		}
		if userInfo.Credit <= 0 {
			return gerror.New("会员余额不足")
		}
		err = User.ChangeCredit(ctx, userInfo.Id, -userInfo.Credit, model.CreditWithdraw)
		return nil
	})

}

//处理门票转账
func (s *listenTask) DealTransfer(c context.Context, param *chainService.BscGameTransferLog) error {
	return dao.FaBscTransfer.Transaction(c, func(ctx context.Context, tx *gdb.TX) error {
		data := &model.FaBscTransfer{
			From:    param.FromAddr.String(),
			To:      param.ToAddr.String(),
			Amount:  BigIntToF(param.Value, model.TokenDecimals),
			Created: int(time.Now().Unix()),
		}
		_, err := dao.FaBscTransfer.Ctx(ctx).OmitEmpty().Save(data)
		if err != nil {
			g.Log().Debug("Service DoListenTask DealTransfer Transfer Save Err:", err)
		}
		return err
	})

}
