// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal



// FaBscTask is the golang structure for table fa_bsc_task.
type FaBscTask struct {
    Id     int    `orm:"id,primary" json:"id"`     //                                
    Type   string `orm:"type"       json:"type"`   // 任务类型                       
    Task   string `orm:"task"       json:"task"`   // 任务内容                       
    Remark string `orm:"remark"     json:"remark"` // 备注                           
    Status int    `orm:"status"     json:"status"` // 状态:0-待完成,1-已完成,2-失败  
}