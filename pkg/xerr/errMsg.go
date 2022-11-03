package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	message[REUQEST_PARAM_ERROR] = "参数错误"
	message[TOKEN_EXPIRE_ERROR] = "token失效，请重新登陆"
	message[TOKEN_GENERATE_ERROR] = "生成token失败"
	message[DB_ERROR] = "数据库繁忙,请稍后再试"
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = "更新数据影响行数为0"
	//用户模块
	message[AccountNotFound] = "账号不存在"
	message[AccountDisable] = "账号不可用"
	message[AccountExist] = "账号已存在"
	message[GetTokenError] = "获取TOKEN失败"
	message[PhoneNumExists] = "手机号已存在"
	message[PhoneNumNotFound] = "手机号不存在"
	message[PassWordWrong] = "密码不正确，请重新输入"
	message[GenTokenFail] = "生成TOKEN失败"
	message[PassWordMust] = "密码必填"
	message[AddressNotFound] = "密码必填"
	//商品模块
	message[ProductNotFound] = "商品不存在"
	message[CategoryNotFound] = "商品分类不存在"

}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
