package xerr

//成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

//全局错误码
const (
	SERVER_COMMON_ERROR           uint32 = 100010
	REUQEST_PARAM_ERROR           uint32 = 100020
	TOKEN_EXPIRE_ERROR            uint32 = 100030
	TOKEN_GENERATE_ERROR          uint32 = 100040
	DB_ERROR                      uint32 = 100050
	DB_UPDATE_AFFECTED_ZERO_ERROR uint32 = 100060
)

//用户模块:20000
const (
	AccountNotFound  uint32 = 200010
	AccountDisable   uint32 = 200020
	AccountExist     uint32 = 200030
	GetTokenError    uint32 = 200040
	PhoneNumExists   uint32 = 200050
	PhoneNumNotFound uint32 = 200060
	PassWordWrong    uint32 = 200070
	GenTokenFail     uint32 = 200080
	PassWordMust     uint32 = 200090
	AddressNotFound  uint32 = 200100
)

//商品服务相关：300000
const (
	ProductNotFound = 3000010
	CategoryNotFound = 3000020
)

