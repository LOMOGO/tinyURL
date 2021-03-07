package errCode

var (
	Success = newError(0, "成功")
	//错误码说明： 第一位为1代表错误类型是系统错误，若是2则代表错误类型为普通错误
	//中间两位错误码代表 引起错误的模块
	//第4 5位错误码代表所处模块的具体错误

	//系统错误前缀为100
	ServerError       = newError(10001, "服务器内部错误")
	GenderURLError    = newError(10002, "短链生成失败")
	GenderQRCodeError = newError(10003, "二维码生成失败")

	//数据库错误的前缀为201
	DatabaseError = newError(20101, "数据库连接错误")

	//用户错误的前缀为202
	NotFound      = newError(20201, "网页迷路了")
	InvalidParams = newError(20202, "入参错误")
)
