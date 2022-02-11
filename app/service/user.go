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

//注册
func (s *user) Register(c context.Context, param *model.UserRegisterReq) error {
	return dao.FaBscUser.Transaction(c, func(ctx context.Context, tx *gdb.TX) error {
		count, err := dao.FaBscUser.Where("address=?", param.UserAddr).LockUpdate().Count()
		if err != nil {
			g.Log().Debug("Service User Register User Count Err:", err)
		}
		if count != 0 {
			return gerror.New("User registered")
		}
		if param.RefAddr != "" {
			if count, _ = dao.FaBscUser.Where("address=?", param.RefAddr).Count(); count == 0 {
				return gerror.New("The referee does not exist")
			}
		} else {
			param.RefAddr = model.OwnAddr
		}
		//注册,添加交易任务
		taskInfo := model.FaBscTask{
			Type:     model.SendRegister,
			Task:     gconv.String(param),
			Created:  int(time.Now().Unix()),
			TaskTime: int(time.Now().Unix()),
		}

		return User.AddTask(ctx, &taskInfo)
	})
}

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

//会员最新活动信息
func (s *user) GetUserGameInfo(ctx context.Context, uid int) (*model.UserGameInfo, error) {
	data := &model.UserGameInfo{}
	userGame, err := dao.FaBscUserGame.Ctx(ctx).Where("uid=?", uid).Order("id desc").One()
	if err != nil || userGame == nil {
		g.Log().Debug("Service User GetUserGameInfo UserGame Find Err:", err)
		return nil, err
	}
	data.FaBscUserGame = userGame
	data.Percent = int((userGame.WillNum - userGame.ReturnNum) * 100 / userGame.InvestNum)
	data.TotalReward, err = dao.FaBscUserGame.Ctx(ctx).Where("uid=? and status>2", uid).Sum(dao.FaBscUserGame.Columns.WillNum)
	if err != nil {
		g.Log().Debug("Service User GetUserGameInfo UserGame Sum Total Err:", err)
		return nil, err
	}
	horse := Index.GetHorseInfo()
	res := Index.GetUserHorse(horse, data.InvestNum)
	data.ImgPath = res.Path
	return data, nil

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
			"ref_credit": &gdb.Counter{
				Field: "ref_credit",
				Value: amount,
			},
			"total_credit": &gdb.Counter{
				Field: "total_credit",
				Value: amount,
			},
			"updated": time.Now().Unix(),
		}
		filename = model.UserCreditFour
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
		err = User.AddTask(ctx, &taskInfo)
		if err != nil {
			g.Log().Debug("Service User ChangeCredit Task Save Err:", err)
			return err
		}
		filename = model.UserCreditThree
	case model.CreditRefWithdraw: //推荐奖提现
		update = g.Map{
			"ref_credit": &gdb.Counter{
				Field: "ref_credit",
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
		err = User.AddTask(ctx, &taskInfo)
		if err != nil {
			g.Log().Debug("Service User ChangeCredit Task Save Err:", err)
			return err
		}
		filename = model.UserCreditFour
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
		err = User.AddTask(ctx, &taskInfo)
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

//收益排行榜
func (s *user) RewardTop(ctx context.Context, req *model.PageReq) (*model.UserRewardTop, error) {
	var (
		data = new(model.UserRewardTop)
		err  error
		sql  = dao.FaBscUser.Ctx(ctx).Order("total_credit desc")
	)
	data.Total, err = sql.Count()
	if err != nil {
		g.Log().Debug("Service User RewardTop Count Err:", err)
		return nil, err
	}
	data.Page = req.Page
	data.Size = req.Size
	err = sql.Page(req.Page, req.Size).Scan(&data.List)
	if err != nil {
		g.Log().Debug("Service User RewardTop Scan Err:", err)
		return nil, err
	}
	if len(data.List) > 0 {
		base := (req.Page - 1) * req.Size
		for k, _ := range data.List {
			data.List[k].Rank = base + k + 1
		}
	}
	return data, err
}

//转账记录
func (s *user) TransferList(ctx context.Context, req *model.UserInfoRep) (*model.UserTransferList, error) {
	var (
		data = new(model.UserTransferList)
		err  error
		sql  = dao.FaBscTransfer.Ctx(ctx).Where("fromAddr=? or toAddr=?", req.UserAddr, req.UserAddr).Order("created desc")
	)
	data.Total, err = sql.Count()
	if err != nil {
		g.Log().Debug("Service User TransferList Count Err:", err)
		return nil, err
	}
	data.Page = req.Page
	data.Size = req.Size
	err = sql.Page(req.Page, req.Size).Scan(&data.List)
	if err != nil {
		g.Log().Debug("Service User TransferList Scan Err:", err)
		return nil, err
	}
	if len(data.List) > 0 {
		for k, v := range data.List {
			if v.FromAddr == req.UserAddr {
				data.List[k].TransferType = 2
			} else {
				data.List[k].TransferType = 1
			}
		}
	}
	return data, err
}

//粮草兑换记录
func (s *user) TicketList(ctx context.Context, req *model.UserInfoRep) (*model.UserTickerList, error) {
	var (
		data = new(model.UserTickerList)
		err  error
		sql  = dao.FaBscUserTicket.Ctx(ctx).Where("uid=?", req.UserId).Order("created desc")
	)
	data.Total, err = sql.Count()
	if err != nil {
		g.Log().Debug("Service User TicketList Count Err:", err)
		return nil, err
	}
	data.Page = req.Page
	data.Size = req.Size
	err = sql.Page(req.Page, req.Size).Scan(&data.List)
	if err != nil {
		g.Log().Debug("Service User TicketList Scan Err:", err)
		return nil, err
	}
	return data, nil
}

//一级直推会员
func (s *user) FirstLevelTeam(ctx context.Context, req *model.UserInfoRep) (*model.TeamUserList, error) {
	var (
		data = new(model.TeamUserList)
		err  error
		sql  = dao.FaBscUser.Ctx(ctx).Where("ref_id=?", req.UserId).Order("created desc")
	)
	data.Total, err = sql.Count()
	if err != nil {
		g.Log().Debug("Service User FirstLevelTeam Count Err:", err)
		return nil, err
	}
	data.Page = req.Page
	data.Size = req.Size
	err = sql.Page(req.Page, req.Size).Scan(&data.List)
	if err != nil {
		g.Log().Debug("Service User FirstLevelTeam Scan Err:", err)
		return nil, err
	}
	if len(data.List) > 0 {
		for k, v := range data.List {
			data.List[k].Num, err = dao.FaBscUserTicket.Ctx(ctx).Where("uid=?", v.Id).Sum(dao.FaBscUserTicket.Columns.PayNum)
			if err != nil {
				g.Log().Debug("Service User FirstLevelTeam UserTicket Sum Err:", err)
			}
		}
	}
	return data, err
}

//会员获取投资回报
func (s *user) GameReward(c context.Context, Uid int) error {
	return dao.FaBscUserGame.Transaction(c, func(ctx context.Context, tx *gdb.TX) error {
		userGameInfo, err := dao.FaBscUserGame.Ctx(ctx).Where("uid=? and status=2", Uid).One()
		if err != nil || userGameInfo == nil {
			g.Log().Debug("Api User GetGameReward userGameInfo Find Err:", err)
			return gerror.New("Information doesn't exist")
		}
		_, err = dao.FaBscUserGame.Ctx(ctx).Where("id=?", userGameInfo.Id).Update(g.Map{"status": 3, "updated": time.Now().Unix()})
		if err != nil {
			g.Log().Debug("Api User GetGameReward userGameInfo Update Err:", err)
			return err
		}
		//发放到账户
		err = s.ChangeCredit(ctx, userGameInfo.Uid, userGameInfo.WillNum, model.CreditReward)
		return err
	})
}

//推荐奖统计
func (s *user) UserRefReward(uid, monLen int) ([]*model.UserRefRewardInfo, error) {
	var err error
	//统计前5个月数据
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	data := make([]*model.UserRefRewardInfo, monLen)

	for i := monLen - 1; i >= 0; i-- {
		data[monLen-i-1] = &model.UserRefRewardInfo{}
		startTime := thisMonth.AddDate(0, -i, 0).Unix()
		endTime := thisMonth.AddDate(0, -i+1, 0).Add(-time.Second).Unix()
		data[monLen-i-1].Month = thisMonth.AddDate(0, -i, 0).Format("01") + "月"
		data[monLen-i-1].TotalReward, err = dao.FaBscCredit.Where("uid=? and type=? and created>=? and created<=?", uid, model.CreditRefReward, startTime, endTime).Sum("num")
	}

	return data, err
}

//发布Bsc任务
func (s *user) AddTask(ctx context.Context, data *model.FaBscTask) error {
	_, err := dao.FaBscTask.Ctx(ctx).OmitEmpty().Save(data)
	if err != nil {
		g.Log().Debug("Service User AddTask Save Err:", err)
		return err
	}
	_, err = g.Redis().DoVar("PUBLISH", "bsc:task", g.Map{"doType": 1})
	if err != nil {
		g.Log().Debug("Service User AddTask Redis Err:", err)
	}
	return err
}
