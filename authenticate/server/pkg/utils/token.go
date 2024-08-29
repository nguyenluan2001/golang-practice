package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nguyenluan2001/golang-authenticate/model"
	"github.com/valyala/fasthttp"
)

const SECRET_KEY = "luannguyen"

func EncodeJWT(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      email,
		"created_at": int32(time.Now().Unix()),
	})
	return token.SignedString([]byte("luannguyen"))
}
func DecodeJWT(tokenStr string) (model.JWTToken, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &model.JWTToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	var tokenRequest model.JWTToken
	if err != nil {
		return tokenRequest, err
	}

	if claims, ok := token.Claims.(*model.JWTToken); ok {
		fmt.Println("claims", claims)
		tokenRequest := model.JWTToken{
			Email:     claims.Email,
			CreatedAt: claims.CreatedAt,
		}

		if len(tokenRequest.Email) == 0 {
			return tokenRequest, errors.New("token was invalid (0)")
		}

		return tokenRequest, nil
	}

	return tokenRequest, errors.New("can not valid JWT token")
}
func SetCookie(ctx *fasthttp.RequestCtx, key, val string, expired int) {
	now := time.Now()
	now = now.Add(time.Second * time.Duration(expired))

	var c fasthttp.Cookie
	c.SetKey(key)
	c.SetValue(val)
	c.SetPath("/")
	c.SetHTTPOnly(true)
	c.SetExpire(now)
	c.SetMaxAge(expired)
	c.SetSameSite(fasthttp.CookieSameSiteLaxMode)
	ctx.Response.Header.SetCookie(&c)
}
func DeleteCookie(ctx *fasthttp.RequestCtx, key string) {
	c := fasthttp.AcquireCookie()
	c.SetKey(key)
	c.SetDomain("localhost")
	c.SetPath("/")
	c.SetExpire(fasthttp.CookieExpireDelete)
	ctx.Response.Header.SetCookie(c)
	fasthttp.ReleaseCookie(c)
}
