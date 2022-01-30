package resps

type ResponseForm struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
}

var (
	// Success 请求成功
	Success = ResponseForm{
		Status: 10000,
		Info:   "success",
	}

	// ReceiveSuccess 请求收到，但不保证成功执行
	// 用于不可靠的请求，比如说大流量时触发了降级，导致无法对所有请求提供可靠服务时
	ReceiveSuccess = ResponseForm{
		Status: 10001,
		Info:   "received",
	}

	// Param 请求参数错误
	Param = ResponseForm{
		Status: 20001,
		Info:   "param error",
	}

	VerifyFailed = ResponseForm{
		Status: 20002,
		Info:   "verify failed",
	}

	// TokenExpired token过期
	TokenExpired = ResponseForm{
		Status: 20003,
		Info:   "token expired",
	}

	// RefreshTokenExpired refresh token过期
	RefreshTokenExpired = ResponseForm{
		Status: 20004,
		Info:   "登录过期，请重新登录",
	}

	// PermissionRefused 没有该操作的权限
	PermissionRefused = ResponseForm{
		Status: 20005,
		Info:   "你没有该操作的权限",
	}

	// UsernameOfPasswordError 登录时账号密码错误
	UsernameOfPasswordError = ResponseForm{
		Status: 20006,
		Info:   "账号或用户名错误",
	}

	// JWZXGG jwzx不可用
	JWZXGG = ResponseForm{
		Status: 30001,
		Info:   "教务在线不可用",
	}

	// JWZXParseFailed 教务在线解析失败
	JWZXParseFailed = ResponseForm{
		Status: 30002,
		Info:   "教务在线解析失败",
	}

	// CenterServiceError 中心服务超时、出现意外错误
	CenterServiceError = ResponseForm{
		Status: 30003,
		Info:   "center server error",
	}

	// ServiceDemoted 触发服务熔断或降级时的返回
	ServiceDemoted = ResponseForm{
		Status: 30004,
		Info:   "该服务暂时不可用",
	}

	// DatabaseUnavailable 数据库不可用
	DatabaseUnavailable = ResponseForm{
		Status: 40001,
		Info:   "database unavailable",
	}

	// CacheUnavailable redis错误
	CacheUnavailable = ResponseForm{
		Status: 40002,
		Info:   "cache unavailable",
	}

	// Internal 内部服务错误，出现panic或者严重的error
	Internal = ResponseForm{
		Status: 50000,
		Info:   "internal server error",
	}
)
