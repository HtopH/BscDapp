// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

// FaBscTransfer is the golang structure for table fa_bsc_transfer.
type FaBscTransfer struct {
	Id      int    `orm:"id,primary" json:"id"`      //
	From    string `orm:"from"       json:"from"`    // 转账地址
	To      string `orm:"to"         json:"to"`      // 接收地址
	Amount  string `orm:"amount"     json:"amount"`  // 数量
	Created int    `orm:"created"    json:"created"` //
}
