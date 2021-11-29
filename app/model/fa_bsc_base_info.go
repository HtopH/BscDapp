// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"BscDapp/app/model/internal"
)

// FaBscBaseInfo is the golang structure for table fa_bsc_base_info.
type FaBscBaseInfo internal.FaBscBaseInfo

// Fill with you ideas below.
//系统基本信息
type BscBaseInfo struct {
	OwnAddr       string  //根节点地址
	TokenDecimal  string  //代币精度
	TicketPercent float64 //门票兑换比例-100%
	SpendTicket   float64 //已兑换门票
	JoinPercent   float64 //参与活动金额与门票比例
}
