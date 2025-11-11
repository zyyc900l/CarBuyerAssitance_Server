package errno

// 错误码的设计原则是, 以尽可能少的错误码来传递必要的信息,
// 让前端能够根据尽量少的 error code 和具体的场景来告知用户错误信息
// 总的来说前端不依赖于后端传递的 msg 来告知用户, 而是通过 code 来额外处理
// 当然如果有一些强指向性错误信息, 你当然可以再写进来一个 code, 比如密码错误或者用户已存在
// 我们将这种与业务强相关的 code 也放在 errno 包中, 主要是为了方便统一管理与避免 code 冲突

// 业务处理成功
const (
	SuccessCode = 10000
	SuccessMsg  = "success"
)

// 参数
const (
	ParamVerifyErrorCode  = 20000 + iota // 提供参数有问题
	ParamMissingErrorCode                //参数缺失
)

// 鉴权
const (
	AuthInvalidCode        = 30000 + iota // 鉴权失败
	AuthAccessExpiredCode                 // 访问令牌过期
	AuthRefreshExpiredCode                // 刷新令牌过期
	AuthPermissionCode                    // 令牌等级不够，如stu无法进行审核
	AuthNoTokenCode                       // 没有 token
)

// 业务错误
const (
	// user
	ServiceUserExistCode     = 40000 + iota
	ServiceUserPasswordError // 密码错误
	ServiceUserNotExistCode  //用户不存在错误
	ServiceUserBanExistCode
	ServiceGiftNotExistCode
	ServiceGiftRunOutExistCode
	ServicePointRunOutExistCode //积分不足
)

// 服务错误
const (
	InternalServiceErrorCode = 50000 + iota // 内部服务错误
	InternalDatabaseErrorCode
)
