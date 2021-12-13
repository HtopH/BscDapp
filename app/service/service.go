package service

import (
	"BscDapp/app/dao"
	"BscDapp/app/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/shopspring/decimal"
	"math"
	"math/big"
	"strconv"
	"time"
)

var (
	baseNum float64 = 100000 //计算门票兑换变化基数
)

type JsonResponse struct {
	Code    int         `json:"code"`    // 错误码
	Data    interface{} `json:"data"`    // 返回数据
	Message string      `json:"message"` // 提示信息
}

func GetConfig(key string) interface{} {
	res, err := dao.FaConfig.Where(dao.FaConfig.Columns.Name, key).Value(dao.FaConfig.Columns.Value)
	if err != nil {
		g.Log().Debug("Service GetConfig Err:", err)
		return ""
	}
	return res

}

func GetGameEndTime(gameInfo *model.FaBscGameInfo, nowSeed float64) int {
	var (
		newEnd int
		now    = time.Now()
	)
	if gameInfo.JackPool < model.GameLevelTwoPool && nowSeed >= model.GameLevelTwoPool {
		newEnd = int(now.Unix() + model.GameLevelTwoTime)
	} else if gameInfo.JackPool < model.GameLevelThreePool && nowSeed >= model.GameLevelThreePool {
		newEnd = int(now.Unix() + model.GameLevelThreeTime)
	} else if gameInfo.JackPool < model.GameLevelFourPool && nowSeed >= model.GameLevelFourPool {
		newEnd = int(now.Unix() + model.GameLevelFourTime)
	} else if gameInfo.JackPool < model.GameLevelFivePool && nowSeed >= model.GameLevelFivePool {
		newEnd = int(now.Unix() + model.GameLevelFiveTime)
	}
	return newEnd
}

func GetBigInt(num float64, decimals string) *big.Int {
	dec := gconv.Float64(decimals) // 兑换比例精度
	temp, _ := decimal.NewFromFloat(num).Mul(decimal.NewFromFloat(dec)).Float64()
	float := strconv.FormatFloat(temp, 'f', -1, 64)
	res, _ := new(big.Int).SetString(float, 10)
	return res
}

func BigIntToF(num *big.Int, decimals string) float64 {
	return gconv.Float64(num.String()) / gconv.Float64(decimals)
}

func GetPercent() (tempN int, percent float64, spend float64) {
	tempN, spend = GetNo()
	base := float64(GetBase(tempN))
	n := float64(tempN)
	//直接算浮点会丢失精度,需要用函数处理浮点和运算
	decimalBase := decimal.NewFromFloat(base)
	if tempN == 1 {
		percent = 1
	} else if n > 1 && n < 11 {
		percent, _ = decimal.NewFromFloat(1).Sub(decimal.NewFromFloat(float64(tempN - 1)).Div(decimalBase)).Float64()
		//percent = 1 - float64(tempN-1)/base
	} else {
		if (tempN-10)%9 > 0 {
			percent, _ = decimal.NewFromFloat(10).Div(decimalBase).Sub(decimal.NewFromFloat(float64((tempN - 10) % 9)).Div(decimalBase)).Float64()
			//percent = 10/base - float64((tempN-10)%9)/base
		} else {
			percent, _ = decimal.NewFromFloat(1).Div(decimalBase).Float64()
			//percent = 1 / base
		}
	}
	return tempN, percent, spend
}

func GetBase(n int) int {
	var base float64 = 10
	if n > 10 {
		b := float64((n-11)/9 + 2)
		base = math.Pow(base, b)
	}
	return int(base)
}

func GetNo() (int, float64) {
	var n int
	spendTickets, err := dao.FaBscBaseInfo.Where(dao.FaBscBaseInfo.Columns.TheKey, model.BaseSpendKey).Value(dao.FaBscBaseInfo.Columns.TheValue)
	if err != nil {
		g.Log().Debug("Service service GetNo BaseInfo Value Err:", err)
		return n, 0
	}
	if spendTickets.Float64() > 0 {
		n = int(math.Ceil(spendTickets.Float64() / baseNum))
	} else {
		n = 1
	}
	return n, spendTickets.Float64()
}
