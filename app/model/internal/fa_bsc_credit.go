// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal



// FaBscCredit is the golang structure for table fa_bsc_credit.
type FaBscCredit struct {
    Id      int     `orm:"id,primary" json:"id"`      //         
    Uid     int     `orm:"uid"        json:"uid"`     // 会员ID  
    Type    string  `orm:"type"       json:"type"`    // 类型    
    Num     float64 `orm:"num"        json:"num"`     // 数量    
    Created string  `orm:"created"    json:"created"` //         
}