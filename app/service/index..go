package service

import (
	"BscDapp/app/dao"
	"BscDapp/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"time"
)

var Index = indexService{}

type indexService struct {
}

//获取系统基本信息
func (s *indexService) GetBaseInfo() *model.BscBaseInfo {
	var data = &model.BscBaseInfo{}
	data.TokenDecimal = model.TokenDecimals
	data.JoinPercent = model.PercentBase / model.PercentJoinTicket
	data.OwnAddr = model.OwnAddr
	data.SpendNum, data.TicketPercent, data.SpendTicket = GetPercent()
	return data
}

//首页投资详情
func (s *indexService) GetGameBaseInfo() (*model.GameIndexInfo, error) {
	data := &model.GameIndexInfo{}
	res, err := dao.FaBscGameInfo.Where("status=1").Order("round desc").One()
	if err != nil {
		g.Log().Debug("Api Index GameInfo GameInfo One Err:", err)
		return nil, err
	}
	//结束时间返回剩余时间
	res.EndTime = res.EndTime - int(time.Now().Unix())

	data.FaBscGameInfo = res
	data.FaBscGameInfo.SeedPool += gconv.Float64(GetConfig("add_seed"))
	data.FaBscGameInfo.JackPool += gconv.Float64(GetConfig("add_reward"))

	data.TotalInvest, err = dao.FaBscUserGame.Where("status=1").Sum(dao.FaBscUserGame.Columns.InvestNum)
	if err != nil {
		g.Log().Debug("Api Index GameInfo UserGame TotalInvest Err:", err)
		return nil, err
	}
	data.TotalInvest += gconv.Float64(GetConfig("add_invest"))

	data.TotalReward, err = dao.FaBscUserGame.Where("status>1").Sum(dao.FaBscUserGame.Columns.InvestNum)
	if err != nil {
		g.Log().Debug("Api Index GameInfo UserGame TotalReward Err:", err)
		return nil, err
	}
	data.TotalReward += gconv.Float64(GetConfig("add_credit"))
	return data, nil
}
