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
