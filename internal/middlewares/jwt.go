package middlewares

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	flog "github.com/xgxw/foundation-go/log"
)

/*
	是否有必要将 jwt 封装到基础组建中: 否

	考虑下那些模块/代码应该复用, 那些不应该. jwt变化较多, 不属于项目间复用的模块.
	jwt组件主要分为三部分:
	1. jwt数据结构与构造函数
	2. jwt处理函数(handle)
	3. jwt数据载体(payload)

	显然 2/3 是不同的. 但其实, 1也是隐藏的不同的. 如某些项目认证需要数据库链接, 有些从缓存取, 有些通过网址认证, 这显然造成1也是不同的. 所以 jwt 不应被封装到基础组件中.
	其次, 为了内聚度更高, 函数应有 HandleFunc, 以便有认证路由时, 直接提供认证函数(因为加密解密需要用一套信息.)
*/

/*
	是否将认证逻辑放入 jwt 中, 还是单独路由

	后端是否有login路由有两种处理逻辑
	当使用cookie存储token时, 可以不需要login路由, 因为认证中间件中可以直接设置cookie.
	当使用localstorage/queryParam时, 需要login路由, 因为需要写入到response.data中.

	当前端有登录页时, 需要login路由, 否则登录页没有合适的请求地址. 并且前端可以不判断 未认证 错误, 因为后端可以直接重定向到登录页. 但是不建议这么做, 因为跳转前最好通知用户确认.
	当不使用登录页时, 可以在页面弹窗提示输入认证信息, 且必须验证后端返回值有无 未认证 错误.
	当 没有login路由&&没有登录页 时, 前端必须在每个请求判断有无认证信息, 有则写入cookie中, 或者将认证信息写入cookie, 但这样不安全.

	目前采用的方案:
	1. 前端添加登录页, for 某些进入页面就需要权限的网页. 所以, 对应的, 后端也需要 auth 路由
	2. 对于前端某些提交时才需要权限的网页, 添加密钥输入提示浮层.

	对于后端, 显然 jwt 组件是该项目的一部分而不是共用的一部分. 业务处理逻辑与数据载体都是耦合与项目的
*/

const (
	// DefaultExpires 默认存活时间
	DefaultExpires = time.Hour * 2
)

type (
	// AuthenticationOptions 认证中间件配置项
	AuthenticationOptions struct {
		Key     string `json:"-" yaml:"key" mapstructure:"key"`
		Expires int64  `json:"-" yaml:"expires" mapstructure:"expires"`
		Cipher  string `json:"-" yaml:"cipher" mapstructure:"cipher"`
	}
	// JWTMiddleware : only use HMAC
	JWTMiddleware struct {
		key           []byte
		signingMethod jwt.SigningMethod
		parse         *jwt.Parser
		logger        *flog.Logger
		// 用于claims
		expires time.Duration
		// 用于认证, 后续可根据需要改为db密码
		cipher string
	}

	payloadClaims struct {
		*jwt.StandardClaims
		UserID uint `json:"user_id" mapstructure:"user_id"`
	}
)

// NewJWTMiddlewares 生成JWT中间件
func NewJWTMiddlewares(logger *flog.Logger, opts AuthenticationOptions) *JWTMiddleware {
	expires := DefaultExpires
	if opts.Expires != 0 {
		expires = time.Duration(opts.Expires)
	}
	jwt := &JWTMiddleware{
		key:           []byte(opts.Key),
		signingMethod: jwt.GetSigningMethod("HS256"),
		parse:         new(jwt.Parser),
		logger:        logger,
		expires:       expires,
		cipher:        opts.Cipher,
	}
	return jwt
}

// MiddlewareFunc 生成中间件函数
func (this *JWTMiddleware) MiddlewareFunc(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		tokenStr := c.Request().Header.Get("Authorization")
		if tokenStr == "" {
			return c.NoContent(http.StatusForbidden)
		}
		payload, err := this.verifyToken(tokenStr)
		if err != nil {
			this.logger.Errorf("verifyToken error: %v", err)
			return this.redirectAuth(c)
		}

		c.Set("user_id", payload.UserID)
		return next(c)
	}
}

func (this *JWTMiddleware) redirectAuth(c echo.Context) error {
	return c.NoContent(http.StatusForbidden)
}

// HandlerFunc 生成路由处理函数, 主要用于请求路由生成token.
func (this *JWTMiddleware) HandlerFunc(c echo.Context) (err error) {
	type requestCarrier struct {
		Cipher string `json:"cipher" form:"cipher" query:"cipher"`
	}
	r := new(requestCarrier)
	if err := c.Bind(r); err != nil {
		return c.NoContent(http.StatusForbidden)
	}
	if r.Cipher != this.cipher {
		return c.NoContent(http.StatusForbidden)
	}
	standClaims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(this.expires).Unix(),
	}
	payload := &payloadClaims{
		UserID:         1,
		StandardClaims: standClaims,
	}
	token, _ := this.signToken(payload)
	type responseCarrier struct {
		Token     string `json:"token"`
		ExpiresAt int64  `json:"expires_at"`
	}
	resp := responseCarrier{
		Token:     token,
		ExpiresAt: standClaims.ExpiresAt,
	}
	return c.JSON(http.StatusOK, resp)
}

func (this *JWTMiddleware) signToken(claims jwt.Claims) (tokenStr string, err error) {
	token := jwt.NewWithClaims(this.signingMethod, claims)
	out, err := token.SignedString(this.key)
	return out, err
}

func (this *JWTMiddleware) verifyToken(tokenStr string) (payload *payloadClaims, err error) {
	token, err := this.parse.ParseWithClaims(tokenStr, &payloadClaims{}, func(t *jwt.Token) (interface{}, error) {
		return this.key, nil
	})
	if err != nil {
		return nil, err
	}
	if payload, ok := token.Claims.(*payloadClaims); ok && token.Valid {
		return payload, nil
	}
	return nil, errors.Wrap(err, "parse jwt token error")
}
