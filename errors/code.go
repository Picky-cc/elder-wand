package errors

type ResponseCode int

type JSONResponse struct {
	Code    ResponseCode
	Message string
	Data    interface{}
}

// 约定：错误码取值范围[1000, MAX_INT]
const (
	Success ResponseCode = 1 // 成功

	// Error Code
	UnknownError              ResponseCode = 1000 // 未知错误
	DBQueryError              ResponseCode = 1001 // 数据库查询错误
	ParamError                ResponseCode = 1002 // 参数错误
	ParamParsingError         ResponseCode = 1003 // 参数解析错误
	ObjectExistError          ResponseCode = 1004 // 对象已存在
	HttpNotStatusOKError      ResponseCode = 1005 // http状态码非200错误
	HttpBusinessCodeError     ResponseCode = 1006 // http请求业务码错误
	NotFoundError             ResponseCode = 1007 // 对象不存在
	ImmutableData             ResponseCode = 1008 // 数据不可被修改
	DBDataError               ResponseCode = 1009 // 数据错误
	ExistAgreementWithAccount ResponseCode = 1010 // 账户已经分配给清算或审计协议
	PaymentError              ResponseCode = 1011 // 支付错误
	FetchFileError            ResponseCode = 1012 // 获取文件错误
	StoreFileError            ResponseCode = 1013 // 存储文件错误
	FileExistError            ResponseCode = 1014 // 文件已存在
	RoutingNotFoundError      ResponseCode = 1015 // 溯源错误
	FileHandlerNotExistError  ResponseCode = 1016 // 文件处理器不存在
	FileExceedMaxSizeError    ResponseCode = 1017 // 文件超过最大限制
	DataSourceError           ResponseCode = 1018 // 数据源错误
	OpenFileError             ResponseCode = 1019 // 打开文件失败
	DuplicateError            ResponseCode = 1020 // 已存在
	UnsupportedError          ResponseCode = 1021 // 不支持错误
	BusinessLogicError        ResponseCode = 1022 // 业务逻辑错误
	DBDataConfigError         ResponseCode = 1023 // 数据库配置表内容错误
	AuthFailedError           ResponseCode = 1024 // 数据库配置表内容错误
	ResponseUnmarshalError    ResponseCode = 1025 // 报文解析错误
	ValueError                ResponseCode = 1026 // 值不合法
	FileReadError             ResponseCode = 1027 // 文件读取失败

	UndefinedError ResponseCode = 9998 // 未定义错误
	IgnoreError    ResponseCode = 9999 // 可忽略不记log错误
)

var codeMessages = map[ResponseCode]string{
	Success:                   "成功",
	UnknownError:              "未知错误",
	DBQueryError:              "数据库查询错误",
	ParamError:                "参数错误",
	ParamParsingError:         "参数解析错误",
	ObjectExistError:          "对象已存在",
	HttpNotStatusOKError:      "http状态码非200错误",
	HttpBusinessCodeError:     "http请求业务码错误，非成功状态1",
	NotFoundError:             "对象不存在",
	ImmutableData:             "数据不可被修改",
	DBDataError:               "数据错误",
	ExistAgreementWithAccount: "账户已经分配给清算或审计协议",
	PaymentError:              "支付错误",
	FetchFileError:            "获取文件错误",
	StoreFileError:            "存储文件错误",
	FileExistError:            "文件已存在",
	RoutingNotFoundError:      "溯源失败",
	FileHandlerNotExistError:  "文件处理器不存在",
	FileExceedMaxSizeError:    "文件超过最大限制",
	DataSourceError:           "数据源错误",
	OpenFileError:             "打开文件失败",
	DuplicateError:            "已存在",
	BusinessLogicError:        "业务逻辑错误",
	DBDataConfigError:         "数据库配置表内容错误",
	AuthFailedError:           "校验失败",
	ResponseUnmarshalError:    "报文解析错误",
	ValueError:                "值不合法",
}

func GetMessage(code ResponseCode) string {
	msg, ok := codeMessages[code]
	if ok {
		return msg
	}
	return string(code)
}
