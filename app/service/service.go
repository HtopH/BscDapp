package service

import (
	"BscDapp/app/model"
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
