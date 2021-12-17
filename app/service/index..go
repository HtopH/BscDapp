package service

import (
	"BscDapp/app/dao"
	"BscDapp/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"strings"
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
	data.SpendNum -= 1
	imgs := GetConfig("imgPath")
	data.ImgInfo = gconv.Map(imgs)
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

	invented := gconv.Map(GetConfig("invented_num"))

	data.FaBscGameInfo = res
	data.FaBscGameInfo.SeedPool += gconv.Float64(invented["add_seed"])
	data.FaBscGameInfo.JackPool += gconv.Float64(invented["add_reward"])

	data.TotalInvest, err = dao.FaBscUserGame.Where("status=1").Sum(dao.FaBscUserGame.Columns.InvestNum)
	if err != nil {
		g.Log().Debug("Api Index GameInfo UserGame TotalInvest Err:", err)
		return nil, err
	}
	data.TotalInvest += gconv.Float64(invented["add_invest"])

	data.TotalReward, err = dao.FaBscUserGame.Where("status>1").Sum(dao.FaBscUserGame.Columns.InvestNum)
	if err != nil {
		g.Log().Debug("Api Index GameInfo UserGame TotalReward Err:", err)
		return nil, err
	}
	data.TotalReward += gconv.Float64(invented["add_credit"])
	return data, nil
}

//获取所以马匹信息
func (s *indexService) GetHorseInfo() []*model.HorseInfo {
	data := make([]*model.HorseInfo, 0)
	info := gconv.Map(GetConfig("horse_info"))
	price := gconv.Map(GetConfig("horse_price"))
	if len(info) > 0 {
		for k, v := range info {
			temp := model.HorseInfo{
				Name:  k,
				Path:  v.(string),
				Price: price[k].(string),
			}
			data = append(data, &temp)
		}

	}
	return data
}

//获取投资金额对应的马
func (s *indexService) GetUserHorse(horse []*model.HorseInfo, investNum float64) *model.HorseInfo {

	var data = &model.HorseInfo{}
	if len(horse) > 0 {
		for _, v := range horse {
			prices := strings.Split(v.Price, "-")
			if len(prices) == 2 {
				if investNum >= gconv.Float64(prices[0]) && investNum <= gconv.Float64(prices[1]) {
					data = v
					break
				}
			}

		}
	}
	return data
}
