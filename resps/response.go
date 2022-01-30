package resps

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Receive(c *gin.Context) {
	c.JSON(http.StatusOK, ReceiveSuccess)
}

func OK(c *gin.Context) {
	c.JSON(http.StatusOK, Success)
}

func OKWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": 10000,
		"info":   "success",
		"data":   data,
	})
}

func InternalError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, Internal)
}

func VerifyError(c *gin.Context) {
	c.JSON(http.StatusForbidden, VerifyFailed)
}

func ParamError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, Param)
}

func NormWithInfo(c *gin.Context, state int, msg string) {
	c.JSON(http.StatusBadRequest, ResponseForm{
		Status: state,
		Info:   msg,
	})
}

func InternalErrorWithInfo(c *gin.Context, state int, msg string) {
	c.JSON(http.StatusInternalServerError, ResponseForm{
		Status: state,
		Info:   msg,
	})
}

func PermissionError(c *gin.Context) {
	c.JSON(http.StatusForbidden, PermissionRefused)
}

func UsernameOrPasswordError(c *gin.Context) {
	c.JSON(http.StatusOK, UsernameOfPasswordError)
}

func HandleError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, ResponseForm{
		Status: 50099,
		Info:   err.Error(),
	})
}

// StdResp 返回标准格式的response
// 如非业务的特殊需求，尽量使用上述给出的几种返回
func StdResp(c *gin.Context, form ResponseForm) {
	c.JSON(http.StatusOK, form)
}
