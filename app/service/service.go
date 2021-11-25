package service

import (
	"BscDapp/app/model"
	"github.com/gogf/gf/util/gconv"
	"math/big"
	"strconv"
	"time"
)

type JsonResponse struct {
	Code    int         `json:"code"`    // 错误码
	Data    interface{} `json:"data"`    // 返回数据
	Message string      `json:"message"` // 提示信息
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
	dec, _ := strconv.Atoi(decimals) // 兑换比例精度
	temp := num * float64(dec)
	float := strconv.FormatFloat(temp, 'f', -1, 64)
	res, _ := new(big.Int).SetString(float, 10)
	return res
}

func BigIntToF(num *big.Int, decimals string) float64 {
	return gconv.Float64(num.String()) / gconv.Float64(decimals)
}
