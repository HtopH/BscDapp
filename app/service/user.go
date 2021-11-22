package service

import (
	"BscDapp/app/dao"
	"BscDapp/app/model"
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"time"
)

var User = user{}

type user struct{}

//修改会员余额
func (s *user) ChangeCredit(ctx context.Context, uid int, amount float64, doType string) error {
	var (
		update = g.Map{}
		err    error
	)
	if doType == model.CreditPool || doType == model.CreditRefReward {
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
	} else if doType == model.CreditBuyTicket {
		update = g.Map{
			"ticket_num": &gdb.Counter{
				Field: "credit",
				Value: amount,
			},
			"updated": time.Now().Unix(),
		}
	} else {
		return gerror.New("Service User ChangeCredit Err:会员余额事件不存在")
	}

	_, err = dao.FaBscUser.Ctx(ctx).Where("id=?", uid).Update(update)
	if err != nil {
		g.Log().Debug("Service User ChangeCredit User Update Err:", err)
		return err
	}
	creditInfo := model.FaBscCredit{
		Uid:     uid,
		Type:    doType,
		Num:     amount,
		Created: int(time.Now().Unix()),
	}
	_, err = dao.FaBscCredit.Ctx(ctx).OmitEmpty().Save(creditInfo)
	if err != nil {
		g.Log().Debug("Service User ChangeCredit Credit Save Err:", err)
	}
	return err
}
