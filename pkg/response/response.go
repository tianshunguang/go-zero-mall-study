package response

import (
	"encoding/json"
	"gitee.com/xiaobu0712/go-zero-mall/pkg/xerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
)

type ResponseSuccessBean struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type NullJson struct{}

func Success(data interface{}) *ResponseSuccessBean {
	return &ResponseSuccessBean{xerr.OK, "OK", data}
}

type ResponseErrorBean struct {
	Code int64           `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
}

func Error(errCode int64, errMsg string, data json.RawMessage) *ResponseErrorBean {
	return &ResponseErrorBean{errCode, errMsg, data}
}

//http返回
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	if err == nil {
		//成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回
		errcode := xerr.SERVER_COMMON_ERROR
		//errmsg := xerr.NewErrMsg(xerr.SERVER_COMMON_ERROR)
		data, _ := json.Marshal(NullJson{})
		causeErr := errors.Cause(err)                // err类型
		if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
			//自定义CodeError
			errcode = uint32(int(e.GetErrCode()))
			//errmsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := int64(gstatus.Code())
				errcode = uint32(int(grpcCode))
				//errmsg = gstatus.Message()
			}
		}
		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)
		httpx.WriteJson(w, http.StatusBadRequest, Error(int64(errcode), "errmsg", data))
	}
}

func ParameterError(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	data, _ := json.Marshal(NullJson{})
	httpx.WriteJson(w, http.StatusBadRequest, Error(int64(xerr.REUQEST_PARAM_ERROR), err.Error(), data))

}
func HttpResultParam(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	httpx.WriteJson(w, http.StatusBadRequest, Error(int64(xerr.REUQEST_PARAM_ERROR), err.Error(), nil))
}
