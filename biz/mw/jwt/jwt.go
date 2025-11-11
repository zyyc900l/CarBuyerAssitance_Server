package jwt

import (
	"CarBuyerAssitance/biz/dal/mysql"
	"CarBuyerAssitance/biz/model/user"
	"CarBuyerAssitance/biz/pack"
	"CarBuyerAssitance/biz/service"
	"CarBuyerAssitance/pkg/constants"
	"CarBuyerAssitance/pkg/errno"
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/jwt"
	"log"
	"time"
)

var (
	identityKey               = constants.ContextUserId
	TypeKey                   = constants.TypeKey
	AccessTokenJwtMiddleware  *jwt.HertzJWTMiddleware
	RefreshTokenJwtMiddleware *jwt.HertzJWTMiddleware
)

func AccessTokenJwt() {
	var err error
	AccessTokenJwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:                       "Video",
		Key:                         []byte(constants.AccessTokenKey),
		Timeout:                     time.Hour,
		MaxRefresh:                  time.Hour,
		WithoutDefaultTokenHeadName: true,
		TokenLookup:                 "header: Access-Token",
		IdentityKey:                 identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*mysql.User); ok {
				id := v.UserId
				var role string
				if v.Status == 1 {
					role = "user"
				} else {
					role = "admin"
				}
				return jwt.MapClaims{
					TypeKey:     role,
					identityKey: id,
				}
			}
			return jwt.MapClaims{}
		},

		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return claims[AccessTokenJwtMiddleware.IdentityKey]
		},

		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			err := errno.NewErrNo(int64(code), message)
			pack.SendFailResponse(c, err)
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.Set("Access-Token", token)
		},

		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginStruct user.LoginRequest
			if err := c.BindAndValidate(&loginStruct); err != nil {
				return nil, err
			}
			userResp, err := LoginCheck(ctx, c, loginStruct.Id, loginStruct.Password)
			if err != nil {
				return nil, err
			}
			return userResp, nil
		},
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
}
func LoginCheck(ctx context.Context, c *app.RequestContext, id, password string) (*mysql.User, error) {
	userResp, err := service.NewUserService(ctx, c).Login(id, password)
	if err != nil {
		return nil, err
	}
	return userResp, nil
}
func RefreshTokenJwt() {
	var err error
	RefreshTokenJwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "video zone",
		Key:         []byte(constants.RefreshTokenKey),
		Timeout:     time.Hour * 72,
		MaxRefresh:  time.Hour * 24,
		IdentityKey: identityKey,
		TokenLookup: "header: Refresh-Token",
		//往令牌中添加的信息
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*mysql.User); ok {
				id := v.UserId
				var role string
				if v.Status == 1 {
					role = "user"
				} else {
					role = "admin"
				}
				return jwt.MapClaims{
					TypeKey:     role,
					identityKey: id,
				}
			}
			return jwt.MapClaims{}
		},
		//从令牌中提取信息
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c) // 是从 JWT 令牌中提取 claims 的函数
			log.Printf("claims: %+v", claims)

			// 检查 claims[identityKey] 是否存在
			userID, exists := claims[identityKey]
			if !exists {
				log.Println("claims['userid'] 不存在")
				return nil
			}
			// 将 userID 存储到上下文中
			c.Set(constants.ContextUserId, userID)
			return userID
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			pack.SendFailResponse(c, errno.NewErrNo(int64(code), message))
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.Set("Refresh-Token", token)
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginStruct user.LoginRequest
			if err := c.BindAndValidate(&loginStruct); err != nil {
				return nil, err
			}
			userResp, err := LoginCheck(ctx, c, loginStruct.Id, loginStruct.Password)
			if err != nil {
				return nil, err
			}
			return userResp, nil
		},
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
}

func IsAccessTokenAvailable(ctx context.Context, c *app.RequestContext, rank int) error {
	claims, err := AccessTokenJwtMiddleware.GetClaimsFromJWT(ctx, c)
	if err != nil {
		return errno.AuthNoToken
	}

	switch v := claims["exp"].(type) { //switch-case处理解析出来的时间类型并与当前时间做比较
	case nil:
		return errno.AuthNoToken
	case float64:
		if int64(v) < AccessTokenJwtMiddleware.TimeFunc().Unix() {
			return errno.AuthAccessExpired
		}
	case json.Number:
		n, err := v.Int64()
		if err != nil {
			return errno.NewErrNo(errno.InternalServiceErrorCode, "Token parse error")
		}
		if n < AccessTokenJwtMiddleware.TimeFunc().Unix() {
			return errno.AuthAccessExpired
		}
	default:
		return errno.NewErrNo(errno.InternalServiceErrorCode, "Token parse error")
	}
	c.Set("JWT_PAYLOAD", claims) //将令牌存入上下文
	identity := AccessTokenJwtMiddleware.IdentityHandler(ctx, c)

	if identity != nil {
		c.Set(AccessTokenJwtMiddleware.IdentityKey, identity) //将用户id解析出存入上下文
	}
	var access int
	if claims[constants.TypeKey] == "admin" {
		access = 3
	} else if claims[constants.TypeKey] == "counselor" {
		access = 2
	} else if claims[constants.TypeKey] == "user" {
		access = 1
	}
	if access < rank {
		return errno.NewErrNo(errno.AuthPermissionCode, "only higher token level access")
	}

	if !AccessTokenJwtMiddleware.Authorizator(identity, ctx, c) { //
		return errno.AuthInvalid
	}

	return nil
}
func IsRefreshTokenAvailable(ctx context.Context, c *app.RequestContext, rank int) error {
	claims, err := RefreshTokenJwtMiddleware.GetClaimsFromJWT(ctx, c)
	if err != nil {
		return errno.AuthNoToken
	}

	switch v := claims["exp"].(type) {
	case nil:
		return errno.AuthNoToken
	case float64:
		if int64(v) < RefreshTokenJwtMiddleware.TimeFunc().Unix() {
			return errno.AuthRefreshExpired
		}
	case json.Number:
		n, err := v.Int64()
		if err != nil {
			return errno.NewErrNo(errno.InternalServiceErrorCode, "Token parse error")
		}
		if n < RefreshTokenJwtMiddleware.TimeFunc().Unix() {
			return errno.AuthRefreshExpired
		}
	default:
		return errno.NewErrNo(errno.InternalServiceErrorCode, "Token parse error")
	}

	c.Set("JWT_PAYLOAD", claims)
	identity := RefreshTokenJwtMiddleware.IdentityHandler(ctx, c)
	if identity != nil {
		c.Set(RefreshTokenJwtMiddleware.IdentityKey, identity)
		hlog.Info("dd")
	}
	var access int
	if claims[constants.TypeKey] == "admin" {
		access = 3
	} else if claims[constants.TypeKey] == "counselor" {
		access = 2
	} else if claims[constants.TypeKey] == "user" {
		access = 1
	}
	if access < rank {
		return errno.NewErrNo(errno.AuthPermissionCode, "only higher token level access")
	}
	if !RefreshTokenJwtMiddleware.Authorizator(identity, ctx, c) {
		return errno.AuthInvalid
	}

	return nil
}

func GenerateAccessToken(c *app.RequestContext) {
	data := service.GetUserIDFromContext(c)
	tokenString, _, _ := AccessTokenJwtMiddleware.TokenGenerator(data)
	c.Header("New-Access-Token", tokenString)
}

func Init() {
	AccessTokenJwt()
	RefreshTokenJwt()
	errInit := AccessTokenJwtMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("AccessTokenJwtMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	errInit = RefreshTokenJwtMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("RefreshTokenJwtMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
}
