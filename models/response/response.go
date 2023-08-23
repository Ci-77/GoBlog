package response

import (
	utils "gvb/utils/pwd"

	"github.com/gin-gonic/gin"
)

// 定义返回值
type Resonse struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}
type ListResponse[T any] struct {
	Count int64 `json:"count"`
	List  T   `json:"list"`
}

const (
	Success = 0
	Error   = 7
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(200, Resonse{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}
func Ok(data any, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}
func OkWithData(data any, c *gin.Context) {
	Result(Success, data, "成功", c)
}
func OkWithMsg(msg string, c *gin.Context) {
	Result(Success, map[string]any{}, msg, c)
}
func Fail(data any, msg string, c *gin.Context) {
	Result(Error, data, msg, c)
}
func FailWithData(data any, c *gin.Context) {
	Result(Error, data, "成功", c)
}
func FailWithMsg(msg string, c *gin.Context) {
	Result(Error, map[string]any{}, msg, c)
}
func OkWith(c *gin.Context) {
	Result(Success, map[string]any{}, "成功", c)
}
func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	Result(Error, map[string]any{}, "未知错误", c)
}
func OkwithList(list any, count int64, c *gin.Context) {
	OkWithData(ListResponse[any]{
		List: list,
		Count: count,
},c)
}
func FailError(err error,obj any,c *gin.Context)  {
	msg:=utils.GetValidMsg(err,obj)
	FailWithMsg(msg,c)
}