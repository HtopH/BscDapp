package service

import (
	"BscDapp/app/dao"
	"BscDapp/app/model"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

var Index = indexService{}

type indexService struct {
}

//获取系统基本信息
func (s *indexService) GetBaseInfo() (*model.BscBaseInfo, error) {
	var data = &model.BscBaseInfo{}
	spendInfo, err := dao.FaBscBaseInfo.Where("theKey=?", model.BaseSpendKey).One()
	if err != nil || spendInfo == nil {
		g.Log().Debug("Service Index GetBaseInfo spendInfo Find err:", err)
		return nil, gerror.New("门票信息查询失败")
	}
	data.TokenDecimal = model.TokenDecimals
	data.JoinPercent = model.PercentBase / model.PercentJoinTicket
	data.OwnAddr = model.OwnAddr
	data.SpendTicket = gconv.Float64(spendInfo.TheValue)
	data.TicketPercent = GetPercent()
	data.SpendNum = GetNo() - 1
	return data, err
}

//首页投资详情
func (s *indexService) GetGameBaseInfo() (*model.GameIndexInfo, error) {
	data := &model.GameIndexInfo{}
	res, err := dao.FaBscGameInfo.Where("status=1").Order("round desc").One()
	if err != nil {
		g.Log().Debug("Api Index GameInfo GameInfo One Err:", err)
		return nil, err
	}
	data.FaBscGameInfo = res
	data.TotalInvest, err = dao.FaBscUserGame.Where("status=1").Sum(dao.FaBscUserGame.Columns.InvestNum)
	if err != nil {
		g.Log().Debug("Api Index GameInfo UserGame TotalInvest Err:", err)
		return nil, err

	}
	data.TotalReward, err = dao.FaBscUserGame.Where("status>1").Sum(dao.FaBscUserGame.Columns.InvestNum)
	if err != nil {
		g.Log().Debug("Api Index GameInfo UserGame TotalReward Err:", err)
		return nil, err

	}
	return data, nil
}
