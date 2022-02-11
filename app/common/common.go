package common

const (
	SuccessMsg = "Operation is successful"
	Failure    = "The operation failure"
	LangZH     = "1"
	LangEN     = "2"

	LanguageKey = "BscLang"
)

var MultiMsg = [2](map[string]string){
	map[string]string{
		"Operation is successful":       "操作成功",
		"The operation failure":         "操作失败",
		"Please enter the user address": "请输入用户地址",
		"User registered":               "用户已注册",
		"The referee does not exist":    "推荐人不存在",
		"Information doesn't exist":     "信息不存在",
	},
	map[string]string{
		"Operation is successful":       "Operation is successful",
		"The operation failure":         "The operation failure",
		"Please enter the user address": "Please enter the user address",
		"User registered":               "User registered",
		"The referee does not exist":    "The referee does not exist",
		"Information doesn't exist":     "Information doesn't exist",
	},
}
